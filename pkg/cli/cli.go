package cli

import (
	"errors"

	"github.com/adityalstkp/go-ngx-config/internal/ast"
	"github.com/adityalstkp/go-ngx-config/internal/parser"
)

func NewNgxConfParser(file string) (*ast.Config, error) {
	parser, err := parser.NewParser(file)
	if err != nil {
		return nil, err
	}

	parsed := parser.Parse()
	if parsed == nil {
		return nil, errors.New("cannot be parsed")
	}

	return parsed, nil
}
