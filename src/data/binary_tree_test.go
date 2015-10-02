package data

import (
	"fmt"
	"testing"
)

type visitor struct{}

func (v *visitor) Visit(node *BinaryTreeNode, depth int) {
	fmt.Printf("value = %v\n", node.Value)
}

func TestBinaryTreeNode(t *testing.T) {
	root := &BinaryTreeNode{Value: &Object{"ROOT"}}
	root.Left = &BinaryTreeNode{Value: &Object{"1"}}
	root.Right = &BinaryTreeNode{Value: &Object{"2"}}
	root.Right.Left = &BinaryTreeNode{Value: &Object{"3"}}
	root.Accept(&visitor{}, 1)
}
