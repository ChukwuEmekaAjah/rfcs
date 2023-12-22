package parser

import (
	"fmt"
	"strconv"
)

// Scanner represents a json string scanner
type Scanner struct {
	source  string
	tokens  []Token
	current uint
	start   uint
	line    uint
	errors  []error
}

// NewScanner creates a new Scanner instance
func NewScanner(str string) Scanner {
	return Scanner{source: str, tokens: make([]Token, 0), current: 0, start: 0, line: 1}
}

// ScanTokens scans the tokens
func (scanner *Scanner) ScanTokens() []Token {

	for !scanner.isAtEnd() {
		scanner.start = scanner.current
		scanner.scanToken()
	}

	scanner.tokens = append(scanner.tokens, Token{tokenType: "EOF", lexeme: "", literal: nil, position: scanner.start})
	return scanner.tokens
}

// isAtEnd tells us if we've reached the end of the source code string
func (scanner Scanner) isAtEnd() bool {
	return scanner.current >= uint(len(scanner.source))
}

// scanToken looks for lexemes inside the source string
func (scanner *Scanner) scanToken() {
	c := scanner.advance()

	switch c {
	case '[':
		scanner.addToken("LEFT_PAREN", string(c))
	case ']':
		scanner.addToken("RIGHT_PAREN", string(c))
	case '{':
		scanner.addToken("LEFT_BRACE", string(c))
	case '}':
		scanner.addToken("RIGHT_BRACE", string(c))
	case ':':
		scanner.addToken("COLON", string(c))
	case ',':
		scanner.addToken("COMMA", string(c))
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		scanner.line++
	case '"':
		scanner.string()
	default:
		if scanner.isDigit(c) {
			scanner.number()
		} else if scanner.isAlpha(scanner.peek(0)) || scanner.isDigit(scanner.peek(0)) {
			scanner.identifier()
		} else {
			scanner.reportError(scanner.line, fmt.Sprintf("Unexpected character. %d", c))
		}
	}
}

// advance returns the next character in the source code
func (scanner *Scanner) advance() byte {
	scanner.current++
	return scanner.source[scanner.current-1]
}

func (scanner *Scanner) addToken(tokenType string, literal interface{}) {
	text := scanner.source[scanner.start:scanner.current]
	scanner.tokens = append(scanner.tokens, Token{tokenType: tokenType, lexeme: text, literal: literal, position: scanner.start})
}

func (scanner *Scanner) match(expected byte) bool {
	if scanner.isAtEnd() {
		return false
	}
	if scanner.source[scanner.current] != byte(expected) {
		return false
	}
	scanner.current++
	return true
}

func (scanner *Scanner) peek(position uint) byte {
	if scanner.isAtEnd() {
		return byte(0)
	}

	return scanner.source[scanner.current+position]
}

func (scanner *Scanner) string() {
	for scanner.peek(0) != '"' && !scanner.isAtEnd() {
		if scanner.peek(0) == '\n' {
			// increment the current line in case we find a new line character
			scanner.line++
		}
		scanner.advance()
	}

	if scanner.isAtEnd() {
		scanner.reportError(scanner.line, "Unterminated string.")
		return
	}
	scanner.advance()
	value := scanner.source[scanner.start+1 : scanner.current-1]
	scanner.addToken("STRING", value)
}

func (scanner Scanner) isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (scanner Scanner) isAlpha(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func (scanner *Scanner) number() {
	for scanner.isDigit(scanner.peek(0)) {
		scanner.advance()
	}

	if scanner.peek(0) == '.' && scanner.isDigit(scanner.peek(1)) {
		scanner.advance()
		for scanner.isDigit(scanner.peek(0)) {
			scanner.advance()
		}
	}

	num, _ := strconv.ParseFloat(scanner.source[scanner.start:scanner.current], 64)
	scanner.addToken("NUMBER", num)
}

func (scanner *Scanner) identifier() {
	for scanner.isAlpha(scanner.peek(0)) {
		scanner.advance()
	}

	text := scanner.source[scanner.start:scanner.current]
	var tokenType string
	if text == "true" || text == "false" {
		tokenType = "BOOLEAN"
	} else if text == "null" {
		tokenType = "NULL"
	} else {
		scanner.reportError(scanner.start, fmt.Sprintf("Unexpected character. %d", text[0]))
		return
	}

	scanner.addToken(tokenType, text)
}

func (scanner *Scanner) reportError(position uint, message string) {
	scanner.errors = append(scanner.errors, fmt.Errorf("Error at %d with message %s", position, message))
}
