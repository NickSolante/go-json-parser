package parser

import (
	"fmt"
	l "go-json-parser/lib/lexer"
	"go-json-parser/lib/token"
	"strconv"
)

type Parser struct {
	lexer *l.Lexer
}

func NewParser(l *l.Lexer) *Parser {
	return &Parser{lexer: l}
}

func (p *Parser) Parse() (interface{}, error) {
	var output interface{}
	var err error

	tok := p.lexer.NextToken()

	output, err = p.ParseToken(tok)

	tok = p.lexer.NextToken()

	if tok.Type != token.EOF {
		err = fmt.Errorf("expected end of input but found %s", tok.Literal)
	}
	return output, err
}

func (p *Parser) ParseToken(tok token.Token) (interface{}, error) {
	var value interface{}
	var err error

	switch tok.Type {
	case token.TRUE:
		value = true
	case token.FALSE:
		value = false
	case token.NULL:
		value = nil
	case token.STRING:
		value = tok.Literal
	case token.NUMBER:
		value, err = strconv.ParseFloat(tok.Literal, 64)
		if err != nil {
			return value, err
		}
	case token.OSQRBRACKET:
		value, err = p.ParseArray(make([]interface{}, 0))
	case token.OBRACKET:
		value, err = p.ParseObject(make(map[string]interface{}))
	case token.EOF:
		err = fmt.Errorf("unexpected end of input")
	default:
		err = fmt.Errorf("unknown token %s", tok.Literal)
	}

	return value, err
}

func (p *Parser) ParseObject(obj map[string]interface{}) (interface{}, error) {
	var err error
	tok := p.lexer.NextToken()

	if tok.Type == token.CBRACKET {
		return obj, err
	}

	for {
		// We expect that the first value is a string
		if tok.Type != token.STRING {
			return obj, fmt.Errorf("expected key but found %s", tok.Literal)
		}

		key := tok.Literal

		tok = p.lexer.NextToken()
		// the 2nd is a colon
		if tok.Type != token.COLON {
			return obj, fmt.Errorf("expected name separate but found %s", tok.Literal)
		}

		// 3rd is a value that will be decided for again
		tok = p.lexer.NextToken()
		value, err := p.ParseToken(tok)
		if err != nil {
			return obj, err
		}

		obj[key] = value

		tok = p.lexer.NextToken()
		// we need to make sure that if there is no error there is a comma after it else we break the loop
		if tok.Type != token.COMMA {
			break
		}

		tok = p.lexer.NextToken()
	}

	if tok.Type != token.CBRACKET {
		err = fmt.Errorf(`expected but } found %s`, tok.Literal)
	}

	return obj, err
}

func (p *Parser) ParseArray(arr []interface{}) (interface{}, error) {
	var err error

	tok := p.lexer.NextToken()

	if tok.Type == token.CSQRBRACKET {
		return arr, err
	}

	for {
		value, err := p.ParseToken(tok)
		if err != nil {
			return arr, err
		}

		arr = append(arr, value)

		tok = p.lexer.NextToken()

		if tok.Type != token.COMMA {
			break
		}

		tok = p.lexer.NextToken()
	}

	if tok.Type != token.CSQRBRACKET {
		err = fmt.Errorf("expected end of input but found %s", tok.Literal)
	}

	return arr, err
}
