package token

type TokenType int

const (
	EOF     TokenType = iota // End of file
	ILLEGAL                  // Represents a character that is not recognized

	LPAREN // (
	RPAREN // )
	QUOTE  // '

	IDENTIFIER // function names, variable names e.g (+, -, foo, define)
	NUMBER
	STRING
)

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, ch rune) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}
