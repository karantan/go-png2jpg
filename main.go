package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"png2jpg/image"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

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
		jpegContent, err := image.ToJpeg(image.Imager{}, imageCotent)
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
