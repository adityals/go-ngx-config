package parser

import (
	"errors"

	"github.com/adityals/go-ngx-config/internal/ast"
	"github.com/adityals/go-ngx-config/internal/parser"
)

func NewNgxConfParser(parserOpts NgxConfParserCliOptions) (*ast.Config, error) {
	parser, err := parser.NewParser(parser.ParserOptions{Filepath: parserOpts.Filepath, ParseInclude: parserOpts.ParseInclude})
	if err != nil {
		return nil, err
	}

	parsedConf := parser.Parse()
	if parsedConf == nil {
		return nil, errors.New("cannot be parsed")
	}

	return parsedConf, nil
}

func NewStringNgxConfParser(confString string, parseInclude bool) (*ast.Config, error) {
	parser := parser.NewStringParser(confString, parseInclude)

	parsedConf := parser.Parse()
	if parsedConf == nil {
		return nil, errors.New("cannot be parsed")
	}

	return parsedConf, nil
}
