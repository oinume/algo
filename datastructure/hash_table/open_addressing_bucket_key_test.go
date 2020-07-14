package hash_table

import (
	"fmt"
	"testing"

	"github.com/oinume/algo/testings"
)

func Test_bucketKey_isEmpty(t *testing.T) {
	empty := &bucketKey{data: emptyKey{}}
	testings.AssertEqual(t, true, empty.isEmpty(), "isEmpty")
}

func Test_bucketKey_HashCode(t *testing.T) {
	tests := map[string]struct {
		key1         interface{}
		key2         interface{}
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
		key1, key2 := &bucketKey{data: tt.key1}, &bucketKey{data: tt.key2}
		testings.AssertEqual(
			t,
			tt.sameHashCode,
			key1.HashCode() == key2.HashCode(),
			fmt.Sprintf("%v:sameHashCode=%v", name, tt.sameHashCode),
		)
	}
}
