package parser

import (
	"strconv"

	"github.com/Nearrivers/2024-day3-aoc/lexer"
	"github.com/Nearrivers/2024-day3-aoc/token"
)

type Parser struct {
	l                *lexer.Lexer
	curToken         token.Token
	peekToken        token.Token
	isNextMulEnabled bool
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:                l,
		isNextMulEnabled: true,
	}

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseLine() int {
	lineResult := 0
	for p.curToken.Type != token.EOF {
		lineResult += p.parseExpression()
		p.nextToken()
	}

	return lineResult
}

func (p *Parser) parseExpression() int {
	switch p.curToken.Type {
	case token.MUL:
		return p.parseMulFunction()
	case token.DO:
		if p.expectParentheses() {
			p.isNextMulEnabled = true
		}
		p.nextToken()
		return p.parseExpression()
	case token.DONT:
		if p.expectParentheses() {
			p.isNextMulEnabled = false
		}
		p.nextToken()
		return p.parseExpression()
	default:
		return 0
	}
}

func (p *Parser) parseMulFunction() int {
	expectedMulFunctionTokens := []token.TokenType{
		token.LPARENT,
		token.INT,
		token.COMMA,
		token.INT,
		token.RPARENT,
	}

	firstNumber, secondNumber := 0, 0
	for _, tok := range expectedMulFunctionTokens {
		if !p.expectPeek(tok) {
			return 0
		}

		if tok == token.INT {
			n, err := strconv.Atoi(p.curToken.Literal)
			if err != nil {
				return 0
			}

			if firstNumber == 0 {
				firstNumber = n
				continue
			}

			secondNumber = n
		}
	}

	if !p.isNextMulEnabled {
		return 0
	}

	return firstNumber * secondNumber
}

func (p *Parser) expectParentheses() bool {
	expectedDoExpressionTokens := []token.TokenType{
		token.LPARENT,
		token.RPARENT,
	}

	for _, tok := range expectedDoExpressionTokens {
		if !p.expectPeek(tok) {
			return false
		}
	}

	return true
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	return false
}
