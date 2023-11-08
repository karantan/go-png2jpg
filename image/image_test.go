package image

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io"
	"os"
	mock_main "png2jpg/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestToJpeg(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock_main.NewMockCoder(ctrl)
	assert := assert.New(t)

	t.Run("successful decode and encode", func(t *testing.T) {
		img, _ := os.Open("../fixtures/gopher.png")
		pngBytes, _ := io.ReadAll(img)
		m.EXPECT().Decode(bytes.NewReader(pngBytes))
		buf := new(bytes.Buffer)
		m.EXPECT().Encode(buf, nil, &jpeg.Options{Quality: 80}).Return(nil)
		ToJpeg(m, pngBytes)
	})

	t.Run("wrong type", func(t *testing.T) {
		pngBytes := []byte("foo")
		_, err := ToJpeg(m, pngBytes)
		assert.Equal(err, fmt.Errorf("unable to convert \"text/plain; charset=utf-8\" to jpeg"))
	})

	t.Run("error decoding", func(t *testing.T) {
		img, _ := os.Open("../fixtures/gopher.png")
		pngBytes, _ := io.ReadAll(img)
		m.EXPECT().Decode(bytes.NewReader(pngBytes)).Return(nil, fmt.Errorf("unable to decode"))
		_, err := ToJpeg(m, pngBytes)
		assert.Equal(err, fmt.Errorf("unable to decode"))
	})

	t.Run("error encoding", func(t *testing.T) {
		img, _ := os.Open("../fixtures/gopher.png")
		pngBytes, _ := io.ReadAll(img)
		m.EXPECT().Decode(bytes.NewReader(pngBytes))
		buf := new(bytes.Buffer)
		m.EXPECT().Encode(buf, nil, &jpeg.Options{Quality: 80}).Return(fmt.Errorf("unable to encode"))
		_, err := ToJpeg(m, pngBytes)
		assert.Equal(err, fmt.Errorf("unable to encode"))
	})

}
