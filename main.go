package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"flag"
)


const (
	maxTokens	int = 256
)

var (
	prompt	string				// Prompt to display in repl
	symbols map[string]Symbol	// Pre-defined symbols
)

type Kind int
const (
	Integral	Kind	= iota	// An integer number
	Operation					// A procedure
	Value						// Stores a value (token-only)
	Constant					// Immutable value
	String						// A long series of interwoven threads
	NIL							// Sentinel values are bad
)

// Represents a symbol
type Symbol struct {
	Kind				// Type of the symbol
}

// Represents a token
type Token struct {
	Kind				// Type of the token
	body	string		
}


// A toy forth-like language
func main() {
	flag.StringVar(&prompt, "p", "» ", "Prompt to display")
	flag.Parse()
	
	reader := bufio.NewReader(os.Stdin)

	symbols = map[string]Symbol {
	}

	// Main loop
	repl:
	for {
		fmt.Print(prompt)

		// Just read a line, TODO - make by-rune reading later
		text, err := reader.ReadString('\n')
		if err == io.EOF {
			break repl
		}

		// Ignore empty lines
		if len(text) < 2 {
			continue repl
		}

		// Strip newline
		text = text[:len(text)-1]

		/* Tokenizing */
		
		tokens, err := tokenize(text)
		if err != nil {
			fmt.Println("err: could not parse -", err)
			continue repl
		}
		
		fmt.Println(tokens)
	}

	fmt.Println("\nGoodbye ☺")
}
