package hash_table

import (
	"fmt"
	"testing"

	"github.com/oinume/algo/internal/assert"
)

func Test_bucketKey_isEmpty(t *testing.T) {
	empty := newEmptyBucketKey[string]()
	assert.AssertEqual(t, true, empty.isEmpty(), "isEmpty")
}

func Test_bucketKey_HashCode(t *testing.T) {
	tests := map[string]struct {
		key1         any
		key2         any
		sameHashCode bool
	}{
		"same": {
			key1:         1,
			key2:         1,
			sameHashCode: true,
		},
		"different": {
			key1:         "a",
			key2:         "b",
			sameHashCode: false,
		},
		"different_value_but_same_hashCode": {
			key1:         "abc",
			key2:         "cba",
			sameHashCode: true,
		},
	}
	for name, tt := range tests {
		key1 := &bucketKey[any]{data: tt.key1, state: bucketStateNormal}
		key2 := &bucketKey[any]{data: tt.key2, state: bucketStateNormal}
		assert.AssertEqual(
			t,
			tt.sameHashCode,
			key1.HashCode() == key2.HashCode(),
			fmt.Sprintf("%v:sameHashCode=%v", name, tt.sameHashCode),
		)
	}
}
