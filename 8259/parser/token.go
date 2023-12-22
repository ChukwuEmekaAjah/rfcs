package parser

import "fmt"

// Token represents a token after lexing the source code
type Token struct {
	tokenType string
	lexeme    string
	literal   interface{}
	position  uint
}

// NewToken creates a new token struct object
func NewToken(tokenType string, lexeme string, literal interface{}, position uint) *Token {
	return &Token{tokenType: tokenType, lexeme: lexeme, literal: literal, position: position}
}

func (t Token) String() string {
	return fmt.Sprintf("%s %s [%v]", t.tokenType, t.literal, t.position)
}
