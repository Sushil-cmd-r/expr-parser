package scanner

import (
	"strings"
	"unicode"

	"github.com/sushil-cmd-r/expr-parser/location"
	"github.com/sushil-cmd-r/expr-parser/token"
)

type Scanner struct {
	rdOffset int
	lineNo   int
	currLine string
	Lines    []string
}

func New(source string) *Scanner {
	source = strings.TrimRightFunc(source, unicode.IsSpace)
	lines := strings.Split(source, "\n")

	return &Scanner{
		rdOffset: 0,
		lineNo:   0,
		Lines:    lines,
	}
}

func (s *Scanner) Next() token.Token {
	s.trimLeft()

	for s.atLineEnd() && !s.atEnd() {
		s.getLine()
		s.trimLeft()
	}

	if s.atLineEnd() {
		loc := location.NewLocation(s.lineNo, s.rdOffset+1)
		return token.New(token.Eof, "End of Input", loc)
	}

	ch := s.advance()
	loc := location.NewLocation(s.lineNo, s.rdOffset)
	switch ch {
	// Scan Operators
	case '+':
		return token.New(token.Plus, "+", loc)
	case '-':
		return token.New(token.Minus, "-", loc)
	case '*':
		return token.New(token.Star, "*", loc)
	case '/':
		return token.New(token.Slash, "/", loc)
	// Numbers
	default:
		if isNum(ch) {
			return s.scanNumber(loc)
		}
		return token.New(token.Illegal, string(ch), loc)
	}
}

func (s *Scanner) scanNumber(loc location.Location) token.Token {
	st := s.rdOffset - 1
	for isNum(s.peek()) {
		s.advance()
	}
	lit := s.currLine[st:s.rdOffset]

	return token.New(token.Number, string(lit), loc)
}

func (s *Scanner) advance() byte {
	if s.atLineEnd() {
		return ' '
	}
	s.rdOffset += 1
	return s.currLine[s.rdOffset-1]
}

func (s *Scanner) peek() byte {
	if s.atLineEnd() {
		return ' '
	}
	return s.currLine[s.rdOffset]
}

func (s *Scanner) trimLeft() {
	for s.rdOffset < len(s.currLine) && s.currLine[s.rdOffset] == ' ' {
		s.rdOffset += 1
	}
}

func (s *Scanner) getLine() {
	s.currLine = s.Lines[s.lineNo]
	s.lineNo += 1
	s.rdOffset = 0
}

func (s *Scanner) atLineEnd() bool {
	return s.rdOffset >= len(s.currLine)
}

func (s *Scanner) atEnd() bool {
	return s.lineNo >= len(s.Lines)
}

func isNum(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
