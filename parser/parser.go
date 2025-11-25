package parser

import (
	ast "github.com/technik97/horizon/ast"
	lx "github.com/technik97/horizon/lexer"
	tk "github.com/technik97/horizon/token"
)

type Parser struct {
	l *lx.Lexer

	currToken tk.Token
	peekToken tk.Token
}

func New(l *lx.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens so curr & peek are set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParserExpr() ast.SExpr {
	switch p.currToken.Type {
	case tk.NUMBER:
		return &ast.Number{Value: p.currToken.Literal}

	case tk.IDENTIFIER:
		return &ast.Symbol{Value: p.currToken.Literal}

	case tk.LPAREN:
		return p.parseList()

	case tk.QUOTE:
		return p.parseQuote()
	}

	return nil
}

func (p *Parser) parseList() ast.SExpr {
	list := &ast.List{Elements: []ast.SExpr{}}

	p.nextToken()

	for p.currToken.Type != tk.RPAREN && p.currToken.Type != tk.EOF {
		elem := p.ParserExpr()

		if elem != nil {
			list.Elements = append(list.Elements, elem)
		}
		p.nextToken()
	}

	return list
}

func (p *Parser) parseQuote() ast.SExpr {
	p.nextToken()

	value := p.ParserExpr()

	return &ast.Quote{Value: value}
}
