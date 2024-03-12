package token

type TokenType int

const (
	Illegal TokenType = iota
	Eof

	Number
	Plus
	Minus
	Star
	Slash
)
