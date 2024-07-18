package lexer

import (
	"github.com/rboddy/monkey-interpreter/token"
)

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
    case '=':
      tok = newToken(token.ASSIGN, l.ch)
    case ''
	}
}

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination. We use a byte, and only support ASCII, but this could be extended with runes and support UTF-8. Pg. 19-20 for ref
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
