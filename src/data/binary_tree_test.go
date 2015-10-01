package data

import (
	"fmt"
	"testing"
)

func TestBinaryTreeNode(t *testing.T) {
	root := &BinaryTreeNode{Value: &Object{"ROOT"}}
	root.Left = &BinaryTreeNode{Value: &Object{"1"}}
	root.Right = &BinaryTreeNode{Value: &Object{"2"}}
	root.Accept(func(value *Object) {
		fmt.Printf("value = %v\n", value)
	})
}
