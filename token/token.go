package token

import (
	"fmt"

	"github.com/sushil-cmd-r/expr-parser/location"
)

type Token struct {
	Type    TokenType
	Literal string
	Loc     location.Location
}

func New(Type TokenType, Literal string, Loc location.Location) Token {
	return Token{
		Type:    Type,
		Literal: Literal,
		Loc:     Loc,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("%d:%d:%s ", t.Loc.Row, t.Loc.Col, t.Literal)
}
