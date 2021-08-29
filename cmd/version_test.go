package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowVersion(t *testing.T) {
	version = "v1.0.0"
	buf := new(bytes.Buffer)
	c := newVersionCmd(buf)
	showVersion(c)

	assert.Equal(t, version+"\n", buf.String())
}
