package parser

import (
	"fmt"

	"github.com/sushil-cmd-r/expr-parser/token"
)

type ParserErr struct {
	parser  *Parser
	errType ErrType
	errTok  token.Token
}

type ErrType int

const (
	UnexpectedTokErr ErrType = iota
	UnknownTokErr
	InvalidTokErr
)

func NewParserErr(parser *Parser, errTok token.Token, errType ErrType) ParserErr {
	return ParserErr{
		parser:  parser,
		errType: errType,
		errTok:  errTok,
	}
}

func (p ParserErr) Error() string {
	errLoc := fmt.Sprintf("file: %s, line: %d ", "<stdin>", p.errTok.Loc.Row)
	errLine := p.parser.sc.Lines[p.errTok.Loc.Row-1]
	errMarker := make([]byte, p.errTok.Loc.Col)

	i := 0
	for i < p.errTok.Loc.Col {
		if i == p.errTok.Loc.Col-1 {
			errMarker = append(errMarker, '^')
			break
		}
		errMarker = append(errMarker, ' ')
		i += 1
	}

	var errMsg string
	if p.errType == UnexpectedTokErr {
		errMsg = fmt.Sprintf("SyntaxError: Unexpected %s", p.errTok.Literal)
	} else if p.errType == UnknownTokErr {
		errMsg = fmt.Sprintf("SyntaxError: Unknown token %s", p.errTok.Literal)
	} else if p.errType == InvalidTokErr {
		errMsg = fmt.Sprintf("SyntaxError: invalid syntax")
	}

	return fmt.Sprintf(" %s\n   %s\n   %s\n %s\n", errLoc, errLine, errMarker, errMsg)
}
