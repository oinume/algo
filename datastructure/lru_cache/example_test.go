package lru_cache_test

import (
	"fmt"

	"github.com/oinume/algo/datastructure/lru_cache"
)

func Example_defaultLRUCache_PutAndGet() {
	cache := lru_cache.NewDefault[int, int](2)
	cache.Put(1, 1)
	cache.Put(2, 2)

	v, _ := cache.Get(1) // References key `1`
	fmt.Println(v)

	cache.Put(3, 3) // This operation evicts key `2`
	_, ok := cache.Get(2)
	fmt.Println(ok)
	// Output:
	// 1
	// false
}
