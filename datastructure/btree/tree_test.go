package btree

import "testing"

func TestTree_Insert(t *testing.T) {
	t.Run("insert leaf as root", func(t *testing.T) {
		tree := NewTree(nil)
		if err := tree.Insert(Item{1, "1"}); err != nil {
			t.Fatalf("Insert must not return err: %v", err)
		}
	})

	t.Run("insert ", func(t *testing.T) {
		tree := NewTree(nil)
		insert(t, tree, Item{10, "10"})
		if err := tree.Insert(Item{5, "5"}); err != nil {
			t.Fatalf("Insert must not return err: %v", err)
		}
	})
}

func insert(t *testing.T, tree *Tree, item Item) {
	if err := tree.Insert(item); err != nil {
		t.Fatalf("Insert must not return err: %v", err)
	}
}
