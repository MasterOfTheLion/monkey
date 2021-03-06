package parser

import (
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}