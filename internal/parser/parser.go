package parser

import (
	"os"
	"path/filepath"

	"github.com/adityals/go-ngx-config/internal/ast"
	"github.com/adityals/go-ngx-config/internal/directive"
	"github.com/adityals/go-ngx-config/internal/lexer"
	"github.com/adityals/go-ngx-config/internal/statement"
	"github.com/adityals/go-ngx-config/internal/token"
)

type Parser struct {
	lexer            *lexer.Lexer
	rootConfig       string
	currentToken     token.Token
	followingToken   token.Token
	statementParsers map[string]func() statement.IDirective
	blockWrappers    map[string]func(*directive.Directive) statement.IDirective
}

func NewParser(filePath string) (*Parser, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	l := lexer.NewLexer(f)
	l.File = filePath

	p := newParserFromLexer(l)
	return p, nil
}

func newParserFromLexer(lexer *lexer.Lexer) *Parser {
	rootConfig, _ := filepath.Split(lexer.File)
	parser := &Parser{
		lexer:      lexer,
		rootConfig: rootConfig,
	}

	// must double scan next token until file can be read
	parser.nextToken()
	parser.nextToken()

	parser.blockWrappers = map[string]func(*directive.Directive) statement.IDirective{
		"http": func(d *directive.Directive) statement.IDirective {
			return parser.wrapHttp(d)
		},
		"server": func(d *directive.Directive) statement.IDirective {
			return parser.wrapServer(d)
		},
		"location": func(d *directive.Directive) statement.IDirective {
			return parser.wrapLocation(d)
		},
	}

	return parser
}

func (p *Parser) Parse() *ast.Config {
	return &ast.Config{
		Filepath: p.lexer.File,
		Block:    p.parseBlock(),
	}
}

func (p *Parser) isCurrTokenEqual(tType token.TokenType) bool {
	return p.currentToken.Type == tType
}

func (p *Parser) parseBlock() *ast.Block {
	ctx := &ast.Block{
		Directives: make([]statement.IDirective, 0),
	}

parsing_loop:
	for {
		switch {
		case p.isCurrTokenEqual(token.EOF) || p.isCurrTokenEqual(token.BlockEnd):
			break parsing_loop
		case p.isCurrTokenEqual(token.Keyword):
			ctx.Directives = append(ctx.Directives, p.parseStatement())
		}
		p.nextToken()
	}

	return ctx
}

// TODO: handle directive
func (p *Parser) parseStatement() statement.IDirective {
	directive := &directive.Directive{
		Name: p.currentToken.Literal,
	}

	if sp, ok := p.statementParsers[directive.Name]; ok {
		return sp()
	}

	for p.nextToken(); p.currentToken.IsParameterEligible(); p.nextToken() {
		directive.Parameters = append(directive.Parameters, p.currentToken.Literal)
	}

	for {
		if p.isCurrTokenEqual(token.Comment) {
			p.nextToken()
		} else {
			break
		}
	}

	if p.isCurrTokenEqual(token.BlockStart) {
		directive.Block = p.parseBlock()
		if bw, ok := p.blockWrappers[directive.Name]; ok {
			return bw(directive)
		}
		return directive
	}

	return directive
}

func (p *Parser) nextToken() {
	p.currentToken = p.followingToken
	p.followingToken = p.lexer.Scan()
}

func (p *Parser) wrapHttp(directive *directive.Directive) *ast.Http {
	h, _ := ast.NewHttp(directive)
	return h
}

func (p *Parser) wrapServer(directive *directive.Directive) *ast.Server {
	s, _ := ast.NewServer(directive)
	return s
}

func (p *Parser) wrapLocation(directive *directive.Directive) *ast.Location {
	l, _ := ast.NewLocation(directive)
	return l
}
