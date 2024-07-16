package token

type tokenType string

type Token struct{
  Type tokenType
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  // Identifiers and literals
  IDENT = "IDENT" // add, foobar, x, y, etc.
  INT = "INT"  // 12345

  // Operators
  ASSIGN = "="
  PLUS = "+"

  // Delimiters
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "["
  RBRACE = "]"

  // Keywords
  FUNCTION = "FUNCTION"
  LET = "LET"
)
