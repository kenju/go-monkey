package ast

import "github.com/kenju/go-monkey/token"

type Node interface {
	// used only for debugging and testing
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program node is the root node of every AST parser produces.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name *Identifier // Name to hold the identifier of the biding
	Value Expression
}
func (ls *LetStatement) statementNode() { }
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}
func (i *Identifier) expressionNode() { }
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type ReturnStatement struct {
	Token token.Token // the 'return' token
	ReturnValue Expression
}
func (rs *ReturnStatement) statementNode() { }
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}