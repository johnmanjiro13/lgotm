package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createDstFile(t *testing.T, b []byte, filename string) {
	t.Helper()

	f, err := os.Create(filepath.Join("testdata", filename))
	assert.NoError(t, err)
	defer f.Close()

	_, err = f.Write(b)
	assert.NoError(t, err)

	t.Skip("created destination file.")
}
