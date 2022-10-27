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

	parsed := parser.Parse()
	if parsed == nil {
		return nil, errors.New("cannot be parsed")
	}

	return parsed, nil
}
