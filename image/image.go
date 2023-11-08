package image

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
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
