package testify_try

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloTestify(t *testing.T) {
	assert.Equal(t, 123, 123, "they should be equal")
	//assert.Equal(t, 123, 124, "they should be equal")
}

func TestHelloTestifyRequire(t *testing.T) {
	_ = errors.New
	f := func() error { return nil }
	//f := func() error { return errors.New("error0") }
	require.NoError(t, f(), "testify require example")
}
