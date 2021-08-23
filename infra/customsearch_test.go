package infra

import (
	"bytes"
	"image"
	"image/png"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "image/jpeg"
	_ "image/png"
)

func TestToPng(t *testing.T) {
	src, err := os.Open("testdata/image.jpg")
	assert.NoError(t, err)

	img, format, err := image.Decode(src)
	assert.NoError(t, err)
	assert.Equal(t, "jpeg", format)

	res, err := toPng(img)
	assert.NoError(t, err)

	buf := new(bytes.Buffer)
	err = png.Encode(buf, res)
	assert.NoError(t, err)

	img, format, err = image.Decode(buf)
	assert.NoError(t, err)
	assert.Equal(t, "png", format)
}
