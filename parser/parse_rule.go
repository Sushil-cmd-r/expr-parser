package parser

import (
	"github.com/sushil-cmd-r/expr-parser/ast"
	"github.com/sushil-cmd-r/expr-parser/token"
)

type (
	Precedence int
	PrefixFn   func() (ast.Expr, error)
	InfixFn    func(ast.Expr) (ast.Expr, error)
)

const (
	_ Precedence = iota
	NONE
	TERM
	FACTOR
	UNARY
)

type ParseRule struct {
	precedence Precedence
	prefix     PrefixFn
	infix      InfixFn
}

func NewParseRule(precedence Precedence, prefix PrefixFn, infix InfixFn) ParseRule {
	return ParseRule{
		precedence: precedence,
		prefix:     prefix,
		infix:      infix,
	}
}

func (p *Parser) registerParseRules() {
	p.parseRules[token.Number] = NewParseRule(NONE, p.primary, nil)
	p.parseRules[token.Plus] = NewParseRule(TERM, nil, p.binary)
	p.parseRules[token.Minus] = NewParseRule(TERM, p.unary, p.binary)
	p.parseRules[token.Star] = NewParseRule(FACTOR, nil, p.binary)
	p.parseRules[token.Slash] = NewParseRule(FACTOR, nil, p.binary)
	p.parseRules[token.Eof] = NewParseRule(NONE, nil, nil)
}
