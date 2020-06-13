package main

import (
)

// Determine the Kind of a word
func whatKind(w string) Kind {
	if len(w) < 1 {
		return NIL
	}

	// Test if this is a known symbol
	if s, ok := symbols[w]; ok {
		return s.Kind
	}
	
	// TODO - test for 0x, 0o, 0b (hex, octal, binary)
	
	// Test for a number constant
	if w[0] >= '0' && w[0] <= '9' {
		return Constant
	}
	
	// Test for a string
	if w[0] == '"' {
		return String
	}

	return Value
}

// Take text and turn it into tokens
func tokenize(text string) ([]Token, error) {
	tokens := make([]Token, 0, maxTokens)

	return tokens, nil
}
