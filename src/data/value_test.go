package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjectString(t *testing.T) {
	o := &Object{1}
	t.Log(o)
}

func TestObjectToInt(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1, (&Object{1}).ToInt())
	assert.Equal(-1, (&Object{"abc"}).ToIntDefault(-1))
}
