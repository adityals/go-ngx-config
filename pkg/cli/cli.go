package cli

import (
	"github.com/adityalstkp/go-ngx-config/internal/ast"
	"github.com/adityalstkp/go-ngx-config/internal/parser"
)

func NewNgxConfParser(file string) (*ast.Config, error) {
	parser, err := parser.NewParser(file)
	if err != nil {
		return nil, err
	}

	return parser.Parse(), nil
}
