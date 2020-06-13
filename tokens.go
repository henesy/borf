package main

import (
	"strings"
	"io"
	"errors"
	"fmt"
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
	reader := strings.NewReader(text)
	
	inString := false
	var builder strings.Builder
	
	// Check if a rune is whitespace
	spacing := func(r rune) bool {
		whitespace := []rune{' ', '\t', '\n'}
		
		for _, w := range whitespace {
			if r == w {
				return true
			}
		}
		
		return false
	}
	
	// Read each rune
	loop:
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break loop
			}
			
			return nil, errors.New(fmt.Sprint("could not read runes - ", err))
		}
		
		// Skip whitespace - close out the builder and append the token if not building a string
		if spacing(r) && !inString {
			continue loop
		}
		
		// Handle string management
		if r == '"' {
			if inString {
				// In a string - close it out and append the token
			} else {
				// Not in a string - reset the builder and append our rune
			}
		}
		
		_, err = builder.WriteRune(r)
		if err != nil {
			return nil, errors.New(fmt.Sprint("could not build token - ", err))
		}
		
		fmt.Printf("» reading = \"%v\"\n", string(r))
	}

	return tokens, nil
}
