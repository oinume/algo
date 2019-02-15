package btree

import (
	"os"
	"testing"
)

func TestTree_Insert(t *testing.T) {
	t.Run("insert leaf as root", func(t *testing.T) {
		tree := NewTree(3)
		tree.Insert(5)
		tree.Insert(1)
		tree.Insert(3)
		tree.Dump(os.Stdout)
	})
}
