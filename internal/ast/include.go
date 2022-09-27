package ast

import (
	"github.com/adityals/go-ngx-config/internal/directive"
	"github.com/adityals/go-ngx-config/internal/statement"
)

type Include struct {
	*directive.Directive
	IncludePath string
	Configs     []*Config
}

func (c *Include) GetDirectives() []statement.IDirective {
	directives := make([]statement.IDirective, 0)
	for _, config := range c.Configs {
		directives = append(directives, config.GetDirectives()...)
	}

	return directives
}

func (c *Include) FindDirectives(directiveName string) []statement.IDirective {
	directives := make([]statement.IDirective, 0)
	for _, config := range c.Configs {
		directives = append(directives, config.FindDirectives(directiveName)...)
	}

	return directives
}

func (i *Include) GetName() string {
	return i.Directive.Name
}
