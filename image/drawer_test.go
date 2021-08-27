package image

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLGTM(t *testing.T) {
	tests := map[string]struct {
		width            uint
		height           uint
		expectedFileName string
	}{
		"width: 400, height: 0": {
			width:            400,
			height:           0,
			expectedFileName: "lgtm400x0.png",
		},
		"width: 0, height: 400": {
			width:            0,
			height:           400,
			expectedFileName: "lgtm0x400.png",
		},
		"width: 300, height: 400": {
			width:            300,
			height:           400,
			expectedFileName: "lgtm300x400.png",
		},
	}

	d := NewDrawer()
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			src, err := os.Open("testdata/image.jpg")
			assert.NoError(t, err)
			defer src.Close()

			res, err := d.LGTM(src, tt.width, tt.height)
			assert.NoError(t, err)

			actual := new(bytes.Buffer)
			_, err = actual.ReadFrom(res)
			assert.NoError(t, err)

			if os.Getenv("IS_CREATE_DST_FILE") == "true" {
				createDstFile(t, actual.Bytes(), tt.expectedFileName)
			}

			expectedFile, err := os.Open(filepath.Join("testdata", tt.expectedFileName))
			assert.NoError(t, err)
			expected := new(bytes.Buffer)
			_, err = expected.ReadFrom(expectedFile)
			assert.NoError(t, err)

			assert.Equal(t, expected.Bytes(), actual.Bytes())
		})
	}
}

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

	if os.Getenv("IS_CREATE_DST_FILE") == "true" {
		createDstFile(t, actual.Bytes(), "lgtm.jpg")
	}

	expectedFile, err := os.Open("testdata/lgtm.jpg")
	assert.NoError(t, err)
	defer expectedFile.Close()

	expected, err := io.ReadAll(expectedFile)
	assert.NoError(t, err)

	assert.Equal(t, expected, actual.Bytes())
}

func createDstFile(t *testing.T, b []byte, filename string) {
	t.Helper()

	f, err := os.Create(filepath.Join("testdata", filename))
	assert.NoError(t, err)
	defer f.Close()

	_, err = f.Write(b)
	assert.NoError(t, err)

	t.Skip("created destination file.")
}
