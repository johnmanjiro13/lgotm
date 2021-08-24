package infra

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io"
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

func TestDrawStringToImage(t *testing.T) {
	file, err := os.Open("testdata/lgtm.jpg")
	assert.NoError(t, err)
	expected, err := io.ReadAll(file)
	assert.NoError(t, err)

	file, err = os.Open("testdata/image.jpg")
	assert.NoError(t, err)
	img, _, err := image.Decode(file)
	assert.NoError(t, err)

	res, err := drawStringToImage(img, "LGTM")
	assert.NoError(t, err)

	actual := new(bytes.Buffer)
	err = jpeg.Encode(actual, res, &jpeg.Options{Quality: 100})
	assert.NoError(t, err)

	assert.Equal(t, expected, actual.Bytes())
}
