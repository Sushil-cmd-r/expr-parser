package parser

import (
	"fmt"

	"github.com/sushil-cmd-r/expr-parser/ast"
	"github.com/sushil-cmd-r/expr-parser/scanner"
	"github.com/sushil-cmd-r/expr-parser/token"
)

type Parser struct {
	sc *scanner.Scanner

	prevTok token.Token
	currTok token.Token
	peekTok token.Token

	parserErr bool

	parseRules map[token.TokenType]ParseRule
}

func New(sc *scanner.Scanner) *Parser {
	p := &Parser{
		sc:         sc,
		parseRules: make(map[token.TokenType]ParseRule),
		parserErr:  false,
	}

	p.registerParseRules()
	p.advance()
	p.advance()
	return p
}

func (p *Parser) Parse() ast.Expr {
	expr := p.parseExpr(NONE)
	if p.parserErr == true {
		return nil
	}
	if p.currTok.Type != token.Eof {
		fmt.Printf("SyntaxError: Unexpected Token %v\n", p.currTok.Literal)
		return nil
	}
	return expr
}

func (p *Parser) parseExpr(precedence Precedence) ast.Expr {
	prefix := p.parseRules[p.currTok.Type].prefix
	if prefix == nil {
		fmt.Printf("SyntaxError: Unexpected Token %v\n", p.currTok.Literal)
		p.parserErr = true
		return nil
	}
	left := prefix()

	for precedence < p.parseRules[p.peekTok.Type].precedence {
		infix := p.parseRules[p.peekTok.Type].infix
		if infix == nil {
			return left
		}
		p.advance()

		left = infix(left)
	}

	return left
}

func (p *Parser) binary(left ast.Expr) ast.Expr {
	operator := p.currTok
	precedence := p.parseRules[operator.Type].precedence

	p.advance()
	right := p.parseExpr(precedence)
	return ast.NewBinary(operator, left, right)
}

func (p *Parser) unary() ast.Expr {
	operator := p.currTok
	p.advance()
	right := p.parseExpr(UNARY)

	return ast.NewUnary(operator, right)
}

func (p *Parser) primary() ast.Expr {
	return ast.NewNumber(p.currTok.Literal)
}

func (p *Parser) advance() {
	p.prevTok = p.currTok
	p.currTok = p.peekTok
	p.peekTok = p.sc.Next()
}
