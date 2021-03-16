package parser

import (
	"github.com/thebashshell/0lang/ast"
	"github.com/thebashshell/0lang/lexer"
	"github.com/thebashshell/0lang/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.TokenType
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	return p
}
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
