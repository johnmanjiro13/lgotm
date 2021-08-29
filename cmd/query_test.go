package cmd

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/johnmanjiro13/lgotm/cmd/mock_cmd"
)

func TestQueryCommand_LGTM(t *testing.T) {
	tests := map[string]struct {
		query            string
		width            uint
		height           uint
		expectedFileName string
	}{
		"width: 400, height: 0": {
			query:            "query",
			width:            400,
			height:           0,
			expectedFileName: "lgtm400x0.png",
		},
		"width: 0, height: 400": {
			query:            "query",
			width:            0,
			height:           400,
			expectedFileName: "lgtm0x400.png",
		},
		"width: 300, height: 400": {
			query:            "query",
			width:            300,
			height:           400,
			expectedFileName: "lgtm300x400.png",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCustomSearchRepo := mock_cmd.NewMockCustomSearchRepository(ctrl)

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			src, err := os.Open("testdata/image.jpg")
			assert.NoError(t, err)
			defer src.Close()

			mockCustomSearchRepo.EXPECT().FindImage(gomock.Any(), tt.query).Return(src, nil)
			c := &queryCommand{search: mockCustomSearchRepo}
			res, err := c.lgtm(context.Background(), tt.query, tt.height, tt.width)
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

func TestInitConfig(t *testing.T) {
	tests := map[string]struct {
		cfgFile  string
		apiKey   string
		engineID string
		isEnv    bool
	}{
		"with cfgFile": {
			cfgFile:  "testdata/config",
			isEnv:    false,
			apiKey:   "api_key",
			engineID: "engine_id",
		},
		"with environment variables": {
			cfgFile:  "",
			isEnv:    true,
			apiKey:   "api_key_env",
			engineID: "engine_id_env",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			cfg := &QueryConfig{}

			if tt.isEnv {
				os.Setenv("API_KEY", tt.apiKey)
				os.Setenv("ENGINE_ID", tt.engineID)
				defer func() {
					os.Unsetenv("API_KEY")
					os.Unsetenv("ENGINE_ID")
				}()
			}

			initConfig(&cobra.Command{}, tt.cfgFile, cfg)
			expected := &QueryConfig{
				APIKey:   tt.apiKey,
				EngineID: tt.engineID,
			}
			assert.Equal(t, expected, cfg)
		})
	}
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
