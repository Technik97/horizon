package lexer

import (
	t "github.com/technik97/horizon/token"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	input        string
	position     int // The current position in input (points to current char)
	readPosition int // The current reading position (points to the char after current one)
	ch           rune
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // initialize the lexer by reading the first char

	return l
}

// reads the next character and advances the position
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // 0 represents EOF in rune
	} else {
		r, size := utf8.DecodeRuneInString(l.input[l.readPosition:])
		l.ch = r
		l.position = l.readPosition
		l.readPosition += size
	}
}

// peekChar looks at the next character without advancing the position
func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}

	r, _ := utf8.DecodeRuneInString(l.input[l.readPosition:])
	return r
}

// Checks if a rune is valid for an IDENTIFIER
func isLetterOrSymbolChar(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch) ||
		ch == '+' || ch == '-' || ch == '*' || ch == '/' ||
		ch == '?' || ch == '!' || ch == '=' || ch == '>' || ch == '<'
}

// reads a sequence of valid identifier characters
func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetterOrSymbolChar(l.ch) {
		l.readChar()
	}

	return l.input[start:l.position]
}

// readNumber reads a sequence of digits
func (l *Lexer) readNumber() string {
	start := l.position
	for unicode.IsDigit(l.ch) {
		l.readChar()
	}

	// Check for a single decimal point
	if l.ch == '.' && unicode.IsDigit(l.peekChar()) {
		l.readChar() // consume '.'
		for unicode.IsDigit(l.ch) {
			l.readChar()
		}
	}

	return l.input[start:l.position]
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) NextToken() t.Token {
	var tok t.Token

	l.skipWhitespace()

	switch l.ch {
	case '(':
		tok = t.NewToken(t.LPAREN, l.ch)
	case ')':
		tok = t.NewToken(t.RPAREN, l.ch)
	case '\'':
		tok = t.NewToken(t.QUOTE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = t.EOF
	default:
		// Handle multi-character tokens
		if isLetterOrSymbolChar(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = t.IDENTIFIER
			return tok
		} else if unicode.IsDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = t.NUMBER
			return tok
		} else {
			tok = t.NewToken(t.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}
