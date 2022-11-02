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

type ParserOptions struct {
	Filepath     string
	ParseInclude bool
}

type Parser struct {
	opts              ParserOptions
	lexer             *lexer.Lexer
	rootConfig        string
	currentToken      token.Token
	followingToken    token.Token
	parsedIncludes    map[string]*ast.Config
	statementParsers  map[string]func() statement.IDirective
	blockWrappers     map[string]func(*directive.Directive) statement.IDirective
	directiveWrappers map[string]func(*directive.Directive) statement.IDirective
}

func NewStringParser(confString string, parseInclude bool) *Parser {
	return newParserFromLexer(lexer.NewStringLexer(confString), ParserOptions{ParseInclude: parseInclude})
}

func NewParser(parserOpts ParserOptions) (*Parser, error) {
	f, err := os.Open(parserOpts.Filepath)
	if err != nil {
		return nil, err
	}

	l := lexer.NewLexer(f)
	l.File = parserOpts.Filepath

	p := newParserFromLexer(l, parserOpts)
	return p, nil
}

func newParserFromLexer(lexer *lexer.Lexer, opts ParserOptions) *Parser {
	rootConfig, _ := filepath.Split(lexer.File)
	parser := &Parser{
		lexer:          lexer,
		rootConfig:     rootConfig,
		opts:           opts,
		parsedIncludes: make(map[string]*ast.Config),
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

	parser.directiveWrappers = map[string]func(*directive.Directive) statement.IDirective{
		"include": func(d *directive.Directive) statement.IDirective {
			return parser.wrapInclude(d)
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

	if p.isCurrTokenEqual(token.Semicolon) {
		if dw, ok := p.directiveWrappers[directive.Name]; ok {
			return dw(directive)
		}
		return directive
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
	h, err := ast.NewHttp(directive)
	if err != nil {
		panic(err)
	}

	return h
}

func (p *Parser) wrapServer(directive *directive.Directive) *ast.Server {
	s, err := ast.NewServer(directive)
	if err != nil {
		panic(err)
	}

	return s
}

func (p *Parser) wrapLocation(directive *directive.Directive) *ast.Location {
	l, err := ast.NewLocation(directive)
	if err != nil {
		panic(err)
	}

	return l
}

func (p *Parser) wrapInclude(directive *directive.Directive) *ast.Include {
	i, err := ast.NewInclude(directive)
	if err != nil {
		panic(err)
	}

	if p.opts.ParseInclude {
		includeFilePath := i.IncludePath

		if !filepath.IsAbs(includeFilePath) {
			includeFilePath = filepath.Join(p.rootConfig, i.IncludePath)
		}

		includePaths, err := filepath.Glob(includeFilePath)
		if err != nil {
			panic(err)
		}

		for _, includePath := range includePaths {
			if conf, ok := p.parsedIncludes[includePath]; ok {
				if conf == nil {
					continue
				}
			} else {
				p.parsedIncludes[includePath] = nil
			}

			parser, err := NewParser(ParserOptions{
				Filepath:     includePath,
				ParseInclude: p.opts.ParseInclude,
			},
			)
			if err != nil {
				panic(err)
			}

			config := parser.Parse()
			p.parsedIncludes[includePath] = config
			i.Configs = append(i.Configs, config)
		}
	}

	return i
}
