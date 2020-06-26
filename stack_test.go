package main

import (
	"testing"
)

// Test several repeated pushes to stack
func TestPush(t *testing.T) {
	stack := MkStack()

	stack, err := stack.Push(Word{Kind: Integral, Value: int64(5)})
	if err != nil {
		t.Fatal("first push failed - ", err)
	}

	stack, err = stack.Push(Word{Kind: Integral, Value: int64(3)})
	if err != nil {
		t.Fatal("second push failed - ", err)
	}

	if size := stack.Size(); size != 2 {
		t.Fatal("size incorrect, expected 2, got: ", size)
	}
}

// Test creation of a new stack
func TestNewStack(t *testing.T) {
	stack := MkStack()

	if size := stack.Size(); size != 0 {
		t.Fatal("stack size expected: 0, got: ", size)
	}

	if stack.Next != nil {
		t.Fatal("next was not nil for HEAD at new stack")
	}

	if stack.Word.Kind != HEAD {
		t.Fatal("new stack's head was not HEAD, was: ", stack.Word.Kind)
	}
}
