package scanner

import (
	"fmt"

	"github.com/sushil-cmd-r/expr-parser/token"
)

type Scanner struct {
	source   []byte
	rdOffset int
}

func New(source []byte) *Scanner {
	return &Scanner{
		source:   source,
		rdOffset: 0,
	}
}

func (s *Scanner) Next() token.Token {
	for !s.isAtEnd() {
		ch := s.advance()
		switch ch {
		// Scan Operators
		case '+':
			return token.New(token.Plus, "+")
		case '-':
			return token.New(token.Minus, "-")
		case '*':
			return token.New(token.Star, "*")
		case '/':
			return token.New(token.Slash, "/")
			// Ignore Whitespaces
		case '\n', ' ', '\t', '\r':
		// Numbers
		default:
			if isNum(ch) {
				return s.scanNumber()
			}
			return token.New(token.Illegal, fmt.Sprintf("unknown token: %s.", string(ch)))

		}
	}
	return token.New(token.Eof, "End of Input")
}

func (s *Scanner) scanNumber() token.Token {
	st := s.rdOffset - 1
	for isNum(s.peek()) {
		s.advance()
	}
	lit := s.source[st:s.rdOffset]

	return token.New(token.Number, string(lit))
}

func (s *Scanner) advance() byte {
	if s.isAtEnd() {
		return ' '
	}
	s.rdOffset += 1
	return s.source[s.rdOffset-1]
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return ' '
	}
	return s.source[s.rdOffset]
}

func (s *Scanner) isAtEnd() bool {
	return s.rdOffset >= len(s.source)
}

func isNum(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
