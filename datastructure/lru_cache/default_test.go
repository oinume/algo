package lru_cache_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/oinume/algo/datastructure/lru_cache"
)

func Test_defaultLRUCache_PutAndGet(t *testing.T) {
	cache := lru_cache.NewDefault(2)
	cache.Put(1, 1)
	cache.Put(2, 2)

	if got, want := cache.Get(1), 1; got != want {
		t.Errorf("got %v but want %v", got, want)
	}

	cache.Put(3, 3) // evicts key 2
	if got, want := cache.Get(2), -1; got != want {
		t.Errorf("got %v but want %v", got, want)
	}

	cache.Put(4, 4) // evicts key 1
	if got, want := cache.Get(1), -1; got != want {
		t.Errorf("got %v but want %v", got, want)
	}

	if got, want := cache.Get(3), 3; got != want {
		t.Errorf("got %v but want %v", got, want)
	}

	if got, want := cache.Get(4), 4; got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func Test_defaultLRUCache_PutAndGet2(t *testing.T) {
	/*
		NewDefault(2)
		put(2, 1)
		put(2, 2)
		get(2)
		put(1, 1)
		put(4, 1)
		get(2)
	*/
	cache := lru_cache.NewDefault(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	if got, want := cache.Get(2), 2; got != want {
		t.Errorf("got %v but want %v", got, want)
	}

	cache.Put(1, 1)
	cache.Put(4, 1)
	if got, want := cache.Get(2), -1; got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func Test_defaultLRUCache_Dump(t *testing.T) {
	cache := lru_cache.NewDefault(2)
	dumper, ok := cache.(lru_cache.Dumper)
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
		t.Errorf("got %q must contain %q", got, value)
	}
}
