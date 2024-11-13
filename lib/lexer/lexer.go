package lexer

import "go-json-parser/lib/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '{':
		tok = newToken(token.OBRACKET, l.ch)
	case '}':
		tok = newToken(token.CBRACKET, l.ch)
	case '[':
		tok = newToken(token.OSQRBRACKET, l.ch)
	case ']':
		tok = newToken(token.CSQRBRACKET, l.ch)
	case '(':
		tok = newToken(token.OPAREN, l.ch)
	case ')':
		tok = newToken(token.CPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '"':
		tok = l.readString()
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		if l.isDigit(l.ch) || l.ch == token.MINUS {
			tok = l.readNumber()
		} else if l.isLiteralName(l.ch) {
			tok = l.readLiteral()
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readString() token.Token {
	var tok token.Token
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == token.QUOTE {
			tok = token.Token{Type: token.STRING, Literal: l.input[position:l.position]}
			break
		}
		if l.ch == 0 {
			if l.input[l.position-1] != token.QUOTE {
				tok = token.Token{Type: token.ILLEGAL, Literal: l.input[position:l.position]}
			}
			break
		}
	}
	return tok
}

func (l *Lexer) isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) isLiteralName(ch byte) bool {
	return ch == 't' || ch == 'f' || ch == 'n'
}

func (l *Lexer) readLiteral() token.Token {
	start := l.position
	var tok token.Token
	switch l.ch {
	case 't':
		for i, c := range token.TRUE[1:] {
			if c != rune(l.peekChar()) {
				return token.Token{Type: token.ILLEGAL, Literal: l.input[start : start+i]}
			}
			l.readChar()
		}
		tok = token.Token{Type: token.TRUE, Literal: "true"}
	case 'f':
		for i, c := range token.FALSE[1:] {
			if c != rune(l.peekChar()) {
				return token.Token{Type: token.ILLEGAL, Literal: l.input[start : start+i]}
			}
			l.readChar()
		}
		tok = token.Token{Type: token.FALSE, Literal: "false"}
	case 'n':
		for i, c := range token.NULL[1:] {
			if c != rune(l.peekChar()) {
				return token.Token{Type: token.ILLEGAL, Literal: string(l.input[start : start+i])}
			}
			l.readChar()
		}
		tok = token.Token{Type: token.NULL, Literal: "null"}
	default:
		return token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
	}
	return tok
}

func (l *Lexer) readNumber() token.Token {
	start := l.position
	for l.isDigit(l.peekChar()) || l.peekChar() == token.DECIMALPOINT || l.peekChar() == token.MINUS {
		l.readChar()
	}
	return token.Token{Type: token.NUMBER, Literal: string(l.input[start:l.readPosition])}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) skipWhitespace() {
	for {
		if l.ch == token.SPACE {
			l.readChar()
			continue
		}
		if l.ch == token.BACKSLASH && (l.peekChar() == 't' || l.peekChar() == 'n' || l.peekChar() == 'r') {
			l.readChar()
			l.readChar()
			continue
		}
		break
	}
}
