package token

type Token struct {
	Type    TokenType
	Literal string
}

func New(Type TokenType, Literal string) Token {
	return Token{
		Type:    Type,
		Literal: Literal,
	}
}
