package token

// TokenType represents a token string
type TokenType string

// Token represents a token
type Token struct {
	Type    TokenType
	Literal string
}
