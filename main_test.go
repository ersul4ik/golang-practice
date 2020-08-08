package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInOutVariable(t *testing.T) {
	assert := require.New(t)

	const in, out = 4, 2
	assert.Equal(t, in, out)
}
