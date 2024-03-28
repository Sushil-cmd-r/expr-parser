package parser

import (
	"github.com/sushil-cmd-r/expr-parser/ast"
	"github.com/sushil-cmd-r/expr-parser/scanner"
	"github.com/sushil-cmd-r/expr-parser/token"
)

type Parser struct {
	sc *scanner.Scanner

	currTok token.Token
	peekTok token.Token

	parseRules map[token.TokenType]ParseRule
}

func New(sc *scanner.Scanner) *Parser {
	p := &Parser{
		sc:         sc,
		parseRules: make(map[token.TokenType]ParseRule),
	}

	p.registerParseRules()
	p.advance()
	p.advance()
	return p
}

func (p *Parser) Parse() (ast.Expr, error) {
	expr, err := p.parseExpr(NONE)
	if err != nil {
		return nil, err
	}

	if p.peekTok.Type != token.Eof {
		return nil, NewParserErr(p, p.peekTok, InvalidTokErr)
	}

	return expr, nil
}

func (p *Parser) parseExpr(precedence Precedence) (ast.Expr, error) {
	if p.currTok.Type == token.Eof {
		return nil, NewParserErr(p, p.currTok, UnexpectedTokErr)
	}

	prefix := p.parseRules[p.currTok.Type].prefix
	if prefix == nil {
		return nil, NewParserErr(p, p.currTok, UnexpectedTokErr)
	}

	left, err := prefix()
	if err != nil {
		return nil, err
	}

	for {
		rule, ok := p.parseRules[p.peekTok.Type]
		if !ok {
			return nil, NewParserErr(p, p.peekTok, UnknownTokErr)
		}

		peekPrec := rule.precedence
		if peekPrec <= precedence {
			break
		}

		infix := rule.infix
		if infix == nil {
			return left, nil
		}
		p.advance()
		left, err = infix(left)
		if err != nil {
			return nil, err
		}
	}

	return left, nil
}

func (p *Parser) binary(left ast.Expr) (ast.Expr, error) {
	operator := p.currTok
	precedence := p.parseRules[operator.Type].precedence

	p.advance()
	right, err := p.parseExpr(precedence)
	if err != nil {
		return nil, err
	}
	return ast.NewBinary(operator, left, right), nil
}

func (p *Parser) unary() (ast.Expr, error) {
	operator := p.currTok
	p.advance()
	right, err := p.parseExpr(UNARY)
	if err != nil {
		return nil, err
	}

	return ast.NewUnary(operator, right), nil
}

func (p *Parser) primary() (ast.Expr, error) {
	return ast.NewNumber(p.currTok.Literal), nil
}

func (p *Parser) advance() {
	p.currTok = p.peekTok
	p.peekTok = p.sc.Next()
}
