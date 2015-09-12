package data

import (
	"testing"
)

func TestObjectString(t *testing.T) {
	o := NewObjectInt(1)
	t.Log(o)
}
