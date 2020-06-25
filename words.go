package main

import (
	"errors"
	"fmt"
sc	"strconv"
)

// Represents a word
type Word struct {
	Kind 						// Type of the symbol
	Value	interface{}			// Value held in the word - procedures have no value
}


func tokens2words(tokens []Token) ([]Word, error) {
	words := make([]Word, 0, len(tokens))
	
	for i, token := range tokens {
		word := Word{
			Kind: token.Kind,
		}
	
		// Populate word.Value
		switch token.Kind {
		case Procedure:
			// A function with side effects
			// Check if defined (maybe store this?)
			w, ok := symbols[token.body]
			if ok {
				word = w
			} else {
				// TODO - some kind of defining for the procedure?
				// Add to known symbols?
				word.Value = func() error {
								return errors.New("unimplemented procedure!")
							}
			}
		
		case String:
			// A text constant
			word.Value = token.body
		
		case Integral:
			// A number (not constant?)
			// A function that returns the value
			n, err := sc.ParseInt(token.body, 10, 64)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("token #%d invalid Integral: %v", i, err))
			}
			
			word.Value = n
			
		case Real:
			// A floating point number
			// A function that returns the value
			n, err := sc.ParseFloat(token.body, 64)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("token #%d invalid REal: %v", i, err))
			}
			
			word.Value = n
		
		default:
			return nil, errors.New(fmt.Sprintf("token #%d is of unknown type", i))
		}
		
		words = append(words, word)
	}
	
	return words, nil
}

