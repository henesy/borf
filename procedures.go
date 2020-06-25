package main

import (
	"errors"
	"fmt"
)

// The / operation
func Divide(stack *Stack) (*Stack, error) {
	if size := stack.Size(); size < 2 {
		return nil, errors.New(fmt.Sprintf("stack is too small to Divide upon; need: 2, got %d", size))
	}

	// The right hand side of the infix is first off the stack
	stkr, right := stack.Pop()
	stkl, left := stkr.Pop()
	stk := stkl

	if left.Kind != right.Kind {
		return nil, errors.New(fmt.Sprintf("types %v and %v must match", left.Kind, right.Kind))
	}

	switch left.Kind {
	case Integral:
		l := left.Value.(int64)
		r := right.Value.(int64)
		word := Word{
			Kind: left.Kind,
			Value: l / r,
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
			Kind: Real,
			Value: l / r,
		}

		proposed, err := stk.Push(word)
		if err != nil {
			return nil, err
		}

		return proposed, nil

	default:
		return nil, errors.New(fmt.Sprintf("types %v and %v cannot be divided together", left.Kind, right.Kind))
	}

	return stack, nil
}

// The * operation
func Multiply(stack *Stack) (*Stack, error) {
	if size := stack.Size(); size < 2 {
		return nil, errors.New(fmt.Sprintf("stack is too small to Multiply upon; need: 2, got %d", size))
	}

	// The right hand side of the infix is first off the stack
	stkr, right := stack.Pop()
	stkl, left := stkr.Pop()
	stk := stkl

	if left.Kind != right.Kind {
		return nil, errors.New(fmt.Sprintf("types %v and %v must match", left.Kind, right.Kind))
	}

	switch left.Kind {
	case Integral:
		l := left.Value.(int64)
		r := right.Value.(int64)
		word := Word{
			Kind: left.Kind,
			Value: l * r,
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
			Kind: Real,
			Value: l * r,
		}

		proposed, err := stk.Push(word)
		if err != nil {
			return nil, err
		}

		return proposed, nil

	default:
		return nil, errors.New(fmt.Sprintf("types %v and %v cannot be multiplied together", left.Kind, right.Kind))
	}

	return stack, nil
}


// The - operation
func Subtract(stack *Stack) (*Stack, error) {
	if size := stack.Size(); size < 2 {
		return nil, errors.New(fmt.Sprintf("stack is too small to Subtract upon; need: 2, got %d", size))
	}

	// The right hand side of the infix is first off the stack
	stkr, right := stack.Pop()
	stkl, left := stkr.Pop()
	stk := stkl

	if left.Kind != right.Kind {
		return nil, errors.New(fmt.Sprintf("types %v and %v must match", left.Kind, right.Kind))
	}

	switch left.Kind {
	case Integral:
		l := left.Value.(int64)
		r := right.Value.(int64)
		word := Word{
			Kind: left.Kind,
			Value: l - r,
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
			Kind: Real,
			Value: l - r,
		}

		proposed, err := stk.Push(word)
		if err != nil {
			return nil, err
		}

		return proposed, nil

	default:
		return nil, errors.New(fmt.Sprintf("types %v and %v cannot be subtracted together", left.Kind, right.Kind))
	}

	return stack, nil
}


// The + operation
func Add(stack *Stack) (*Stack, error) {
	if size := stack.Size(); size < 2 {
		return nil, errors.New(fmt.Sprintf("stack is too small to Add upon; need: 2, got %d", size))
	}

	// The right hand side of the infix is first
	stkr, right := stack.Pop()
	stkl, left := stkr.Pop()
	stk := stkl

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
			Kind: Real,
			Value: l + r,
		}

		proposed, err := stk.Push(word)
		if err != nil {
			return nil, err
		}

		return proposed, nil

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
