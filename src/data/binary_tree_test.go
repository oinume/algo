package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type visitor1 struct {
	visited []int
	t       *testing.T
}

func (v *visitor1) Visit(node *BinaryTreeNode, depth int) {
	v.visited = append(v.visited, node.Value.ToInt())
	v.t.Logf("value = %v\n", node.Value)
}

func TestBinaryTreeNodeTraversePreorder(t *testing.T) {
	root := &BinaryTreeNode{Value: &Object{1}}
	root.Left = &BinaryTreeNode{Value: &Object{2}}
	root.Right = &BinaryTreeNode{Value: &Object{3}}
	root.Right.Left = &BinaryTreeNode{Value: &Object{100}}
	v1 := &visitor1{
		visited: make([]int, 0, 4),
		t:       t,
	}
	root.TraversePreorder(v1, 1)

	assert := assert.New(t)
	assert.Equal([]int{1, 2, 3, 100}, v1.visited)
}
