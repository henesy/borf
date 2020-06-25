package main

import (
	"fmt"
	"errors"
)

// Represents a stack
type Stack struct {
	Next	*Stack	// Next element down the stack
	Word			// Our {value: type} pair
}

// Create a new stack
func MkStack() *Stack {
	return &Stack{
				Next: nil,
				Word: Word{Kind: HEAD},
			}
}

// Internal pop function - use correctly ☺
func (s *Stack) Pop() (*Stack, Word) {	
	if s.Next == nil {
		// We are the head
		return s, s.Word
	}
	
	return s.Next, s.Word
}

// Pushes to the stack, may shift stack pointer - TODO - locks?
func (s *Stack) Push(w Word) (*Stack, error) {
	switch w.Kind {
	// These are all plain values
	case Integral:
		fallthrough
	case Real:
		fallthrough
	case String:
		top := &Stack{
				Next: s, 
				Word: w,
			}
			
		return top, nil
	
	case Procedure:
		// Call the procedure to modify the stack
		push := w.Value.(func(*Stack) (*Stack, error))
		stk, err := push(s)
		if err != nil {
			return nil, err
		}

		return stk, nil
	
	case Variable:
		// TODO
		return nil, errors.New("variable unimplemented for push")
	case NIL:
		return nil, errors.New("word type of NIL can't be pushed")
	default:
		return nil, errors.New(fmt.Sprint("unknown type: ", w.Kind))
	}
	
	// Change nothing by default?
	return s, nil
}

// Descend to calculate the size of the stack - O(n)
func (s *Stack) Size() uint64 {
	var size uint64 = 0
	
	if s.Word.Kind == HEAD {
		size = 0
	} else {
		size = 1
	}	
	
	if s.Next != nil {
		size += s.Next.Size()
	}
	
	return size
}