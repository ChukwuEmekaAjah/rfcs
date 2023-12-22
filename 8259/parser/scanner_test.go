package parser

import (
	"reflect"
	"testing"
)

func TestNewScanner(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want Scanner
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScanner(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScanner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_ScanTokens(t *testing.T) {
	type fields struct {
		source  string
		tokens  []Token
		current uint
		start   uint
		line    uint
	}
	tests := []struct {
		name   string
		fields fields
		want   []Token
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := &Scanner{
				source:  tt.fields.source,
				tokens:  tt.fields.tokens,
				current: tt.fields.current,
				start:   tt.fields.start,
				line:    tt.fields.line,
			}
			if got := scanner.ScanTokens(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scanner.ScanTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_isAtEnd(t *testing.T) {
	type fields struct {
		source  string
		tokens  []Token
		current uint
		start   uint
		line    uint
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := Scanner{
				source:  tt.fields.source,
				tokens:  tt.fields.tokens,
				current: tt.fields.current,
				start:   tt.fields.start,
				line:    tt.fields.line,
			}
			if got := scanner.isAtEnd(); got != tt.want {
				t.Errorf("Scanner.isAtEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_scanToken(t *testing.T) {
	type fields struct {
		source  string
		tokens  []Token
		current uint
		start   uint
		line    uint
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := &Scanner{
				source:  tt.fields.source,
				tokens:  tt.fields.tokens,
				current: tt.fields.current,
				start:   tt.fields.start,
				line:    tt.fields.line,
			}
			scanner.scanToken()
		})
	}
}

func TestScanner_advance(t *testing.T) {
	type fields struct {
		source  string
		tokens  []Token
		current uint
		start   uint
		line    uint
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := &Scanner{
				source:  tt.fields.source,
				tokens:  tt.fields.tokens,
				current: tt.fields.current,
				start:   tt.fields.start,
				line:    tt.fields.line,
			}
			if got := scanner.advance(); got != tt.want {
				t.Errorf("Scanner.advance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_addToken(t *testing.T) {
	type fields struct {
		source  string
		tokens  []Token
		current uint
		start   uint
		line    uint
	}
	type args struct {
		tokenType string
		literal   interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := &Scanner{
				source:  tt.fields.source,
				tokens:  tt.fields.tokens,
				current: tt.fields.current,
				start:   tt.fields.start,
				line:    tt.fields.line,
			}
			scanner.addToken(tt.args.tokenType, tt.args.literal)
		})
	}
}

func TestScanner_match(t *testing.T) {
	type fields struct {
		source  string
		tokens  []Token
		current uint
		start   uint
		line    uint
	}
	type args struct {
		expected byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := &Scanner{
				source:  tt.fields.source,
				tokens:  tt.fields.tokens,
				current: tt.fields.current,
				start:   tt.fields.start,
				line:    tt.fields.line,
			}
			if got := scanner.match(tt.args.expected); got != tt.want {
				t.Errorf("Scanner.match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_peek(t *testing.T) {
	type fields struct {
		source  string
		tokens  []Token
		current uint
		start   uint
		line    uint
	}
	type args struct {
		position uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := &Scanner{
				source:  tt.fields.source,
				tokens:  tt.fields.tokens,
				current: tt.fields.current,
				start:   tt.fields.start,
				line:    tt.fields.line,
			}
			if got := scanner.peek(tt.args.position); got != tt.want {
				t.Errorf("Scanner.peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_string(t *testing.T) {
	type fields struct {
		source  string
		tokens  []Token
		current uint
		start   uint
		line    uint
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := &Scanner{
				source:  tt.fields.source,
				tokens:  tt.fields.tokens,
				current: tt.fields.current,
				start:   tt.fields.start,
				line:    tt.fields.line,
			}
			scanner.string()
		})
	}
}

func TestScanner_isDigit(t *testing.T) {
	type fields struct {
		source  string
		tokens  []Token
		current uint
		start   uint
		line    uint
	}
	type args struct {
		ch byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := Scanner{
				source:  tt.fields.source,
				tokens:  tt.fields.tokens,
				current: tt.fields.current,
				start:   tt.fields.start,
				line:    tt.fields.line,
			}
			if got := scanner.isDigit(tt.args.ch); got != tt.want {
				t.Errorf("Scanner.isDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_isAlpha(t *testing.T) {
	type fields struct {
		source  string
		tokens  []Token
		current uint
		start   uint
		line    uint
	}
	type args struct {
		ch byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := Scanner{
				source:  tt.fields.source,
				tokens:  tt.fields.tokens,
				current: tt.fields.current,
				start:   tt.fields.start,
				line:    tt.fields.line,
			}
			if got := scanner.isAlpha(tt.args.ch); got != tt.want {
				t.Errorf("Scanner.isAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_number(t *testing.T) {
	type fields struct {
		source  string
		tokens  []Token
		current uint
		start   uint
		line    uint
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := &Scanner{
				source:  tt.fields.source,
				tokens:  tt.fields.tokens,
				current: tt.fields.current,
				start:   tt.fields.start,
				line:    tt.fields.line,
			}
			scanner.number()
		})
	}
}

func TestScanner_identifier(t *testing.T) {
	type fields struct {
		source  string
		tokens  []Token
		current uint
		start   uint
		line    uint
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := &Scanner{
				source:  tt.fields.source,
				tokens:  tt.fields.tokens,
				current: tt.fields.current,
				start:   tt.fields.start,
				line:    tt.fields.line,
			}
			scanner.identifier()
		})
	}
}
