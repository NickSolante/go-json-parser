package main

import (
	"bufio"
	"fmt"
	"go-json-parser/lib/lexer"
	"go-json-parser/lib/parser"
	v "go-json-parser/lib/valid"
	"os"
	"strings"
)

func main() {
	var input string
	var err error

	if len(os.Args) > 1 {
		input = strings.Join(os.Args[1:], " ")
	} else {
		fmt.Println("JSON Parser")
		fmt.Println("Enter JSON input:")

		reader := bufio.NewReader(os.Stdin)
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
	}

	input = strings.TrimSpace(input)

	l := lexer.NewLexer(input)
	p := parser.NewParser(l)

	result, err := p.Parse()
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		fmt.Printf("is JSON Valid? %t\n", v.IsValid(result, err))
		return
	}

	fmt.Printf("Parsed JSON: %+v\n", result)
	fmt.Printf("is JSON Valid? %t\n", v.IsValid(result, err))
}
