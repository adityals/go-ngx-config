package ast

import (
	"github.com/adityalstkp/go-ngx-config/internal/statement"
)

type Config struct {
	*Block
	Filepath string
}

func (n *Config) FindDirectives(directiveName string) []statement.IDirective {
	return n.Block.FindDirectives(directiveName)
}
