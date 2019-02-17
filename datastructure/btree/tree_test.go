package btree

import (
	"fmt"
	"os"
	"testing"
)

func TestTree_Insert(t *testing.T) {
	t.Run("insert leaf as root", func(t *testing.T) {
		// TODO: たぶんsplitChildしたときにおかしくなっている雰囲気なので調査
		tree := NewTree(2)
		for i := 0; i < 4; i++ {
			tree.Insert(int64(i))
			fmt.Printf("--- %d ---\n", i)
			tree.Dump(os.Stdout)
		}
		//tree.Insert(3)
		println("--- final ---")
		tree.Dump(os.Stdout)
	})
}
