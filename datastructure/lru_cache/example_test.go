package lru_cache_test

import (
	"fmt"

	"github.com/oinume/algo/datastructure/lru_cache"
)

func Example_defaultLRUCache_PutAndGet() {
	cache := lru_cache.NewDefault(2)
	cache.Put(1, 1)
	cache.Put(2, 2)

	fmt.Println(cache.Get(1)) // References key `1`

	cache.Put(3, 3) // This operation evicts key `2`
	// Output:
	// 1
	// -1
	fmt.Println(cache.Get(2))
}
