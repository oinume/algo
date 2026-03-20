package lru_cache_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/oinume/algo/datastructure/lru_cache"
)

func Test_defaultLRUCache_PutAndGet(t *testing.T) {
	cache := lru_cache.NewDefault[int, int](2)
	cache.Put(1, 1)
	cache.Put(2, 2)

	if got, ok := cache.Get(1); !ok || got != 1 {
		t.Errorf("got %v, ok %v but want 1, true", got, ok)
	}

	cache.Put(3, 3) // evicts key 2
	if _, ok := cache.Get(2); ok {
		t.Errorf("expected key 2 to be evicted")
	}

	cache.Put(4, 4) // evicts key 1
	if _, ok := cache.Get(1); ok {
		t.Errorf("expected key 1 to be evicted")
	}

	if got, ok := cache.Get(3); !ok || got != 3 {
		t.Errorf("got %v, ok %v but want 3, true", got, ok)
	}

	if got, ok := cache.Get(4); !ok || got != 4 {
		t.Errorf("got %v, ok %v but want 4, true", got, ok)
	}
}

func Test_defaultLRUCache_PutAndGet2(t *testing.T) {
	cache := lru_cache.NewDefault[int, int](2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	if got, ok := cache.Get(2); !ok || got != 2 {
		t.Errorf("got %v, ok %v but want 2, true", got, ok)
	}

	cache.Put(1, 1)
	cache.Put(4, 1)
	if _, ok := cache.Get(2); ok {
		t.Errorf("expected key 2 to be evicted")
	}
}

func Test_defaultLRUCache_StringKeys(t *testing.T) {
	cache := lru_cache.NewDefault[string, int](2)
	cache.Put("a", 1)
	cache.Put("b", 2)

	if got, ok := cache.Get("a"); !ok || got != 1 {
		t.Errorf("got %v, ok %v but want 1, true", got, ok)
	}

	cache.Put("c", 3) // evicts "b"
	if _, ok := cache.Get("b"); ok {
		t.Errorf("expected key 'b' to be evicted")
	}

	if got, ok := cache.Get("c"); !ok || got != 3 {
		t.Errorf("got %v, ok %v but want 3, true", got, ok)
	}
}

func Test_defaultLRUCache_Dump(t *testing.T) {
	cache := lru_cache.NewDefault[int, int](2)
	dumper, ok := any(cache).(lru_cache.Dumper)
	if !ok {
		t.Fatal("Must implement Dumper")
	}

	const value = 12345
	cache.Put(1, value)

	var b bytes.Buffer
	if err := dumper.Dump(&b); err != nil {
		t.Fatalf("Dump failed: %v", err)
	}
	got := b.String()
	if !strings.Contains(got, fmt.Sprint(value)) {
		t.Errorf("got %q must contain %d", got, value)
	}
}
