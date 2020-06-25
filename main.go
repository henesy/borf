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
	symbols map[string]Word		// Pre-defined symbols
)

// Abstract types
type Kind int
const (
	Integral	Kind	= iota	// An integer number
	Real						// A floating point number
	Procedure					// A procedure
	String						// A long series of interwoven threads
	Variable					// A named variable
	NIL							// Sentinel values are bad
	HEAD						// ^
)

// Represents a token
type Token struct {
	Kind				// Type of the token
	body	string
}


// A toy forth-like language
func main() {
	var stack	*Stack				// The stack, for the machine

	flag.StringVar(&prompt, "p", "» ", "Prompt to display")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	// Built-in symbols
	symbols = map[string]Word {
		"+":	Word{Procedure, Add},
		"-":	Word{Procedure, Subtract},
		"*":	Word{Procedure, Multiply},
		"/":	Word{Procedure, Divide},
		".":	Word{Procedure, Pop},
		"pop":	Word{Procedure, Pop},
		// TODO - more builtins
	}

	stack = MkStack()

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

		// Special repl commands
		switch text {
		case "size":
			fmt.Println(stack.Size())
			continue repl

		case "peek":
			fmt.Println(stack.Word)
			continue repl

		case "stack":
			fallthrough
		case "show":
			fmt.Print(stack)
			continue repl

		case "bye":
			break repl

		default:
			;
		}

		tokens, err := tokenize(text)
		if err != nil {
			fmt.Println("err: could not parse -", err)
			continue repl
		}

		//fmt.Println("Tokens =", tokens)

		/* Convert tokens to full - operable - words */

		words, err := tokens2words(tokens)

		//fmt.Println("Words =", symbols)

		/* Push tokens on to stack */

		for i, word := range words {
			proposed, err := stack.Push(word)
			if err != nil {
				fmt.Printf("err: could not push word #%d: %v - %v\n", i, word, err)
				continue repl
			}

			stack = proposed

			fmt.Println("ok.")
		}
	}

	fmt.Println("\nGoodbye ☺")
}
