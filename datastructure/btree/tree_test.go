package btree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTree_Insert_NoSplit(t *testing.T) {
	tests := []struct {
		minDegree int
		input     []int64
		want      []int64
	}{
		{
			minDegree: 2,
			input:     []int64{10, 3, 4},
			want:      []int64{3, 4, 10},
		},
		{
			minDegree: 3,
			input:     []int64{11, 20, 57, 1, 32},
			want:      []int64{1, 11, 20, 32, 57},
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("minDegree:%v", test.minDegree), func(t *testing.T) {
			tree := NewTree(test.minDegree)
			for _, v := range test.input {
				tree.Insert(v)
			}
			if got, want := len(tree.root.keys), len(test.want); got != want {
				t.Errorf("unexpected keys length: got=%v, want=%v", got, want)
			}
			if got, want := tree.root.keys, test.want; !reflect.DeepEqual(got, want) {
				t.Errorf("unexpected keys: got=%v, want=%v", got, want)
			}
		})
	}
}

func TestTree_Insert_Split(t *testing.T) {
	tree := NewTree(2)
	for i := 0; i < 4; i++ {
		tree.Insert(int64(i))
	}
	if got, want := tree.root.keys[0], int64(1); got != want {
		t.Errorf("unexpected root keys[0]: got=%v, want=%v", got, want)
	}
	if got, want := len(tree.root.children), 2; got != want {
		t.Errorf("unexpected root children length: got=%v, want=%v", got, want)
	}
}

func TestTree_Find(t *testing.T) {
	tree := NewTree(2)
	for i := 0; i < 4; i++ {
		tree.Insert(int64(i))
	}
	//tree.Dump(os.Stdout)
	_, index, err := tree.Find(3)
	if err != nil {
		t.Fatalf("tree.Find must succeed: %v", err)
	}
	if got, want := index, 1; got != want {
		t.Errorf("unexpected index returned: got=%v, want=%v", got, want)
	}
}
