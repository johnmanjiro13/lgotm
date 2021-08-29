package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowVersion(t *testing.T) {
	version = "v1.0.0"
	buf := new(bytes.Buffer)
	c := newVersionCmd(buf)
	showVersion(c)

	expected := fmt.Sprintf("lgotm version %s\n", version)
	assert.Equal(t, expected, buf.String())
}
