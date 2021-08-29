package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileCommand_LGTM(t *testing.T) {
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

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := &fileCommand{}
			res, err := c.lgtm("testdata/image.jpg", tt.width, tt.height)
			assert.NoError(t, err)

			actual := new(bytes.Buffer)
			_, err = actual.ReadFrom(res)
			assert.NoError(t, err)

			if os.Getenv("IS_CREATE_DST_FILE") == "true" {
				createDstFile(t, actual.Bytes(), tt.expectedFileName)
			}

			file, err := os.Open(filepath.Join("testdata", tt.expectedFileName))
			assert.NoError(t, err)
			defer file.Close()

			expected := new(bytes.Buffer)
			_, err = expected.ReadFrom(file)
			assert.NoError(t, err)

			assert.Equal(t, expected.Bytes(), actual.Bytes())
		})
	}
}
