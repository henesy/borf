package main

import (
	"errors"
	"fmt"
)


// The + operation
func Add(stack *Stack) (*Stack, error) {
	if size := stack.Size(); size < 2 {
		return nil, errors.New(fmt.Sprintf("stack is too small to Add upon; need: 2, got %d", size))
	}
	
	stkl, left := stack.Pop()
	stkr, right := stkl.Pop()
	stk := stkr
	
	if left.Kind != right.Kind {
		return nil, errors.New(fmt.Sprintf("types %v and %v must match", left.Kind, right.Kind))
	}
	
	switch left.Kind {
	case Integral:
		l := left.Value.(int64)
		r := right.Value.(int64)
		word := Word{
			Kind: left.Kind,
			Value: l + r,
		}
		
		proposed, err := stk.Push(word)
		if err != nil {
			return nil, err
		}
		
		return proposed, nil
		
	case Real:
		l := left.Value.(float64)
		r := right.Value.(float64)
		word := Word{
			Kind: Integral,
			Value: l + r,
		}
		
		proposed, err := stack.Push(word)
		if err != nil {
			return nil, err
		}
		
		stack = proposed
		
		return stack, nil
	
	default:
		return nil, errors.New(fmt.Sprintf("types %v and %v cannot be added together", left.Kind, right.Kind))
	}

	return stack, nil
}

// Pop and display the top of the stack
func Pop(stack *Stack) (*Stack, error) {
	if stack.Next == nil {
		return nil, errors.New("stack underflow → stack is empty")
	}

	stk, word := stack.Pop()
	
	fmt.Println(word.Value)
	
	return stk, nil
}
