package lexer

import "github.com/Nearrivers/2024-day3-aoc/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case 'm':
		tok.Literal = l.readIdentifier()
		tok.Type = token.LookupKeyword(tok.Literal)
		return tok
	case '(':
		tok = newToken(token.LPARENT, l.ch)
	case ')':
		tok = newToken(token.RPARENT, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case 'd':
		if l.peekChar() != 'o' {
			l.readChar()
			tok = newToken(token.ILLEGAL, l.ch)
			return tok
		}

		l.readChar()
		if l.peekChar() != 'n' {
			l.readChar()
			tok = newToken(token.DO, l.ch)
			return tok
		}

		l.readChar()
		if l.peekChar() != '\'' {
			l.readChar()
			tok = token.Token{Type: token.ILLEGAL, Literal: "don"}
			return tok
		}

		l.readChar()
		if l.peekChar() != 't' {
			tok = token.Token{Type: token.ILLEGAL, Literal: "don'"}
			return tok
		}

		tok = token.Token{Type: token.DONT, Literal: "don't"}
		l.readChar()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}

		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '\''
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
