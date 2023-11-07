package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Coder interface used for DI
type Coder interface {
	Decode(r io.Reader) (image.Image, error)
	Encode(w io.Writer, m image.Image, o *jpeg.Options) error
}

// Imager struc type that implements Coder interface
type Imager struct{}

// Decode is a wrapper for png.Decode
func (i Imager) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

// Encode is a wrapper for jpeg.Encode
func (i Imager) Encode(w io.Writer, m image.Image, o *jpeg.Options) error {
	return jpeg.Encode(w, m, o)
}

// ToJpeg converts a PNG image to JPEG format
func ToJpeg(coder Coder, imageBytes []byte) ([]byte, error) {
	contentType := http.DetectContentType(imageBytes)

	switch contentType {
	case "image/png":
		img, err := coder.Decode(bytes.NewReader(imageBytes))

		if err != nil {
			return nil, err
		}

		buf := new(bytes.Buffer)

		if err := coder.Encode(buf, img, &jpeg.Options{Quality: 80}); err != nil {
			return nil, err
		}

		return buf.Bytes(), nil
	}

	return nil, fmt.Errorf("unable to convert %#v to jpeg", contentType)
}

type ReqBody struct {
	URL       string `json:"url" binding:"required"`
	ImageName string `json:"image_name" binding:"required"`
}

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "running on port")
	flag.Parse()
	r := gin.Default()
	r.POST("/convert", func(c *gin.Context) {
		body := ReqBody{}
		if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
			c.JSON(503, gin.H{
				"error": err.Error(),
			})
			return
		}

		out, err := os.CreateTemp("", body.ImageName)
		defer os.Remove(out.Name())
		if err != nil {
			c.JSON(503, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer out.Close()

		resp, err := http.Get(body.URL)
		if err != nil {
			c.JSON(503, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer resp.Body.Close()

		imageCotent, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(503, gin.H{
				"error": err.Error(),
			})
			return
		}
		jpegContent, err := ToJpeg(Imager{}, imageCotent)
		if err != nil {
			c.JSON(503, gin.H{
				"error": err.Error(),
			})
			return
		}

		_, err = out.Write(jpegContent)
		if err != nil {
			c.JSON(503, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.FileAttachment(out.Name(), body.ImageName)
		return
	})
	r.Run(fmt.Sprintf("0.0.0.0:%v", port))
}
