package main

import (
	"bufio"
	"fmt"
	"go-json-parser/lib/lexer"
	"go-json-parser/lib/parser"
	"os"
	"strings"
)

func main() {
	var input string
	var err error

	if len(os.Args) > 1 {
		// Join all arguments after the program name
		input = strings.Join(os.Args[1:], " ")
	} else {
		// Fallback to stdin if no arguments provided
		fmt.Println("JSON Parser")
		fmt.Println("Enter JSON input:")

		reader := bufio.NewReader(os.Stdin)
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
	}

	// Clean input
	input = strings.TrimSpace(input)

	// Initialize the lexer and parser
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)

	result, err := p.Parse()
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Printf("Parsed JSON: %+v\n", result)
}
