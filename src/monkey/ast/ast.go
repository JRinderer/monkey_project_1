package ast

import (
	"go/ast"
	"monkey/token"
	"text/template/parse"
)

type Node interface{
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type LetStatement struct {
	Token token.Token
	Name *parse.Identifier
	Value Exepression
}

func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode(){

}

func (i *Identifier) TokenLiteral() string{
	return i.Token.Literal
}

type Exepression interface{
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string{
	if len(p.Statements)>0 {
		return p.Statements[0].TokenLiteral()
	}else{
		return ""
	}
}
