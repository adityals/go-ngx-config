package parser

import (
	"errors"

	"github.com/adityals/go-ngx-config/internal/ast"
	"github.com/adityals/go-ngx-config/internal/parser"
)

func NewNgxConfParser(parserOpts NgxConfParserCliOptions) (*ast.Config, error) {
	parser, err := parser.NewParser(parserOpts.Filepath)
	if err != nil {
		return nil, err
	}

	parsedConf := parser.Parse()
	if parsedConf == nil {
		return nil, errors.New("cannot be parsed")
	}

	return parsedConf, nil
}

func NewStringNgxConfParser(confString string) (*ast.Config, error) {
	parser := parser.NewStringParser(confString)

	parsedConf := parser.Parse()
	if parsedConf == nil {
		return nil, errors.New("cannot be parsed")
	}

	return parsedConf, nil
}
