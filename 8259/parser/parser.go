package parser

import (
	"fmt"
	"reflect"
)

/*
We will use a recursive descent parser to parse the tokens
The expression for parsing the string is as follows
value = object | array | string | number | null | boolean
object = {(key:value)(, key:value)*}
array = [value(, value)*]
key = string
string = STRING
number = NUMBER
boolean = TRUE | FALSE
null = NULL
*/

// First we need to create a way to loop through the tokens

// Parser reads through tokens and converts it to a map object
type Parser struct {
	tokens  []Token
	current uint
}

// NewParser creates a new instance of parser
func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

// Decode turns json string tokens into a Golang interface
func (p *Parser) Decode() interface{} {
	v := p.value()
	if p.tokens[p.current].tokenType != "EOF" {
		panic(fmt.Sprintf("Invalid characters in json at position %d", p.tokens[p.current].position))
	}
	return v
}

// value tries to parse the tokens in the json string
func (p *Parser) value() interface{} {
	return p.object()
}

// object tries to parse the tokens into a map
func (p *Parser) object() interface{} {
	if p.peek(0).tokenType == "LEFT_BRACE" {
		p.advance()
		v := make(map[string]interface{}, 0)
		k := p.key()
		if k == '_' {
			p.consume("RIGHT_BRACE", "Expected '}' after json object")
			return v
		}
		p.consume("COLON", "Expected ':' after json object key")
		val := p.value()

		v[k.(string)] = val
		for p.match("COMMA") {
			k = p.key()
			p.consume("COLON", "Expected ':' after json object key")
			val = p.value()
			v[k.(string)] = val
		}
		p.consume("RIGHT_BRACE", "Expected '}' after json object")
		return v
	}
	return p.array()
}

// key returns a string
func (p *Parser) key() interface{} {
	k := p.literal()
	if reflect.TypeOf(k).Kind() == reflect.String {
		return k
	}

	if k == '_' {
		return k
	}

	panic(fmt.Sprintf("Expected object key of type string but got %v", k))
}

// array tries to get all the array items
func (p *Parser) array() interface{} {
	if p.peek(0).tokenType == "LEFT_PAREN" {
		p.advance()
		v := make([]interface{}, 0)
		element := p.value()
		if element == '_' {
			p.consume("RIGHT_PAREN", "Expected ']' at the end of array")
			return v
		}
		v = append(v, element)
		for p.match("COMMA") {
			element = p.value()
			v = append(v, element)
		}
		p.consume("RIGHT_PAREN", "Expected ']' at the end of array")
		return v
	}
	return p.literal()
}

func (p *Parser) consume(tokenType string, message string) {
	if p.isAtEnd() {
		panic(message)
	}

	if p.tokens[p.current].tokenType != tokenType {
		panic(fmt.Sprintf("%s but got %v", message, p.tokens[p.current].literal))
	}
	p.advance()
}

// literal returns the base values of a json object
func (p *Parser) literal() interface{} {
	if p.match("BOOLEAN") {
		if p.previous().literal == "true" {
			return true
		}
		return false
	} else if p.match("NUMBER") {
		return p.previous().literal
	} else if p.match("STRING") {
		return p.previous().literal
	} else if p.match("NULL") {
		return nil
	}
	// empty byte sentinel
	return '_'

}

// match checks if the current token matches any of the possible types
func (p *Parser) match(values ...string) bool {
	for _, v := range values {
		if p.peek(0).tokenType == v {
			p.advance()
			return true
		}
	}
	return false
}

// advance moves the tokens counter a step forward
func (p *Parser) advance() {
	if p.isAtEnd() {
		return
	}
	p.current++
}

// previous returns the last seen token
func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}

// isAtEnd tells us if we have reached the end of tokens available
func (p *Parser) isAtEnd() bool {
	return p.peek(0).tokenType == "EOF" || p.peek(0) == Token{}
}

// peek looks ahead to see what character is at the specified position
func (p *Parser) peek(position int) Token {
	if p.current+uint(position) >= uint(len(p.tokens)) {
		return Token{}
	}
	return p.tokens[p.current+uint(position)]
}
