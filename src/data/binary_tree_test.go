package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//       1
//    ／    ＼
//   2       3
//         ／
//       100

type visitor1 struct {
	visited []int
	t       *testing.T
}

func (v *visitor1) Visit(node *BinaryTreeNode, depth int) {
	v.visited = append(v.visited, node.Value.ToInt())
	v.t.Logf("value = %v\n", node.Value)
}

func TestBinaryTreeNodeTraversePreorder(t *testing.T) {
	root := createBinaryTree()
	v1 := &visitor1{
		visited: make([]int, 0, 4),
		t:       t,
	}
	root.TraversePreorder(v1, 1)

	assert := assert.New(t)
	assert.Equal([]int{1, 2, 3, 100}, v1.visited)
}

type visitor2 struct {
	visited []int
	t       *testing.T
}

func (v *visitor2) Visit(node *BinaryTreeNode, depth int) {
	v.visited = append(v.visited, node.Value.ToInt())
	v.t.Logf("value = %v\n", node.Value)
}

func TestBinaryTreeNodeTraverseInorder(t *testing.T) {
	root := createBinaryTree()
	v2 := &visitor2{
		visited: make([]int, 0, 4),
		t:       t,
	}
	root.TraverseInorder(v2, 1)

	assert := assert.New(t)
	assert.Equal([]int{2, 1, 100, 3}, v2.visited)
}

type visitor3 struct {
	visited []int
	t       *testing.T
}

func (v *visitor3) Visit(node *BinaryTreeNode, depth int) {
	v.visited = append(v.visited, node.Value.ToInt())
	v.t.Logf("value = %v\n", node.Value)
}

func TestBinaryTreeNodeTraversePostorder(t *testing.T) {
	root := createBinaryTree()
	v3 := &visitor3{
		visited: make([]int, 0, 4),
		t:       t,
	}
	root.TraversePostorder(v3, 1)

	assert := assert.New(t)
	assert.Equal([]int{2, 100, 3, 1}, v3.visited)
}

func createBinaryTree() *BinaryTreeNode {
	root := &BinaryTreeNode{Value: &Object{1}}
	root.Left = &BinaryTreeNode{Value: &Object{2}}
	root.Right = &BinaryTreeNode{Value: &Object{3}}
	root.Right.Left = &BinaryTreeNode{Value: &Object{100}}
	return root
}
