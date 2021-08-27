package image

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPng(t *testing.T) {
	src, err := os.Open("testdata/image.jpg")
	assert.NoError(t, err)
	defer src.Close()

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
	expectedFile, err := os.Open("testdata/lgtm.jpg")
	assert.NoError(t, err)
	defer expectedFile.Close()

	expected, err := io.ReadAll(expectedFile)
	assert.NoError(t, err)

	file, err := os.Open("testdata/image.jpg")
	assert.NoError(t, err)
	defer file.Close()

	img, format, err := image.Decode(file)
	assert.NoError(t, err)
	assert.Equal(t, "jpeg", format)
	res, err := drawStringToImage(img, "LGTM")
	assert.NoError(t, err)

	actual := new(bytes.Buffer)
	err = jpeg.Encode(actual, res, &jpeg.Options{Quality: 100})
	assert.NoError(t, err)

	assert.Equal(t, expected, actual.Bytes())
}
