package stack

import (
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := NewStack(10)
	stack.Push(1)
	stack.Push(2)
	if got, want := stack.Size(), 2; got != want {
		t.Errorf("something wrong for Push. unexpected stack size: %v", got)
	}
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack(10)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	popped, err := stack.Pop()
	if err != nil {
		t.Fatal("unexpected error from Pop: ", err)
	}
	if want := 3; popped != want {
		t.Errorf("unexpected result from Pop: got=%v, want=%v", popped, want)
	}

	stack.Clear()
	if got, want := stack.Size(), 0; got != want {
		t.Fatalf("unexpected stack size after Clear: got=%v, want=%v", got, want)
	}
}
