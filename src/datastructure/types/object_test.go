package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjectString(t *testing.T) {
	o := &Object{1}
	t.Log(o)
}

func TestObjectInt(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1, (&Object{1}).Int())
	assert.Equal(-1, (&Object{"abc"}).IntDefault(-1))
}
