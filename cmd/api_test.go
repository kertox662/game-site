package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPICmd(t *testing.T) {
	assert.Equal(t, "api", apiCmd.Use)
}
