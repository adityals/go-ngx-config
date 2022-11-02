package ast

import (
	"errors"

	"github.com/adityals/go-ngx-config/internal/directive"
	"github.com/adityals/go-ngx-config/internal/statement"
)

type Include struct {
	*directive.Directive
	IncludePath string
	Configs     []*Config
}

func NewInclude(d *directive.Directive) (*Include, error) {
	if d.Block != nil {
		return nil, errors.New("include cannot have a block")
	}

	if len(d.Parameters) != 1 {
		return nil, errors.New("include must have 1 parameter")
	}

	return &Include{
		Directive:   d,
		IncludePath: d.Parameters[0],
	}, nil
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
