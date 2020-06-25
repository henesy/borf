package main

import (
	"strings"
	"io"
	"errors"
	"fmt"
)

// Determine the Kind of a word
func whatKind(w string) (Kind, error) {
	if len(w) < 1 {
		return NIL, errors.New(fmt.Sprint("empty word"))
	}

	// Test if this is a known symbol
	if s, ok := symbols[w]; ok {
		return s.Kind, nil
	}
	
	// TODO - test for 0x, 0o, 0b (hex, octal, binary)
	
	// Test for a number constant
	if w[0] >= '0' && w[0] <= '9' {
		if strings.Contains(w, "0x") {
			// Hexadecimal (base 16)
			// TODO
		}
		
		if strings.Contains(w, "0o") {
			// Octal (base 8)
			// TODO
		}
		
		if strings.Contains(w, "0b") {
			// Binary (base 2)
			// TODO
		}
	
		if strings.Contains(w, ".") {
			return Real, nil
		}
		
		return Integral, nil
	}
	
	// Test for a string
	if w[0] == '"' {
		return String, nil
	}

	return NIL, errors.New(fmt.Sprintf(`could not determine type of "%s"`, w))
}

// Take a string and return a token
func string2token(s string) (Token, error) {
	var t Token
	t.body = s
	
	kind, err := whatKind(s)
	if err != nil {
		return t, err
	}
	
	t.Kind = kind
	
	return t, nil
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
			if builder.Len() > 0 {
				str := builder.String()
				
				token, err := string2token(str)
				if err != nil {
					return nil, err
				}
				
				tokens = append(tokens, token)
			}
			
			builder.Reset()
		
			continue loop
		}
		
		// Handle string management
		if r == '"' {
			if inString {
				// In a string - close it out and append the token
				builder.WriteRune(r)
				str := builder.String()
				
				token, err := string2token(str)
				if err != nil {
					return nil, err
				}
				
				// We know it's a string
				token.Kind = String
				
				tokens = append(tokens, token)
				
				inString = false
				
				builder.Reset()
				
				continue loop
				
			} else {
				// Not in a string - build a new string token - reset the builder
				builder.Reset()
				builder.WriteRune(r)
				
				inString = true
				
				continue loop
			}
		}
		
		_, err = builder.WriteRune(r)
		if err != nil {
			return nil, errors.New(fmt.Sprint("could not build token - ", err))
		}
		
		//fmt.Printf("» reading = \"%v\"\n", string(r))
	}
	
	// Catch trailing word, will never be a string
	if builder.Len() > 0 {
		str := builder.String()
		
		token, err := string2token(str)
		if err != nil {
			return nil, err
		}
		
		tokens = append(tokens, token)
	}

	return tokens, nil
}
