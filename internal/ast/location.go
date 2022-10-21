package ast

import (
	"errors"

	"github.com/adityals/go-ngx-config/internal/directive"
	"github.com/adityals/go-ngx-config/internal/statement"
)

type Location struct {
	Name       string
	Modifier   string
	Match      string
	Directives []statement.IDirective
}

func NewLocation(directive *directive.Directive) (*Location, error) {
	location := &Location{
		Name:       "location",
		Directives: []statement.IDirective{},
	}

	if block := directive.GetBlock(); block != nil {
		location.Directives = append(location.Directives, block.GetDirectives()...)

		if len(directive.Parameters) == 0 {
			return nil, errors.New("not enough argument in location block")
		}

		if len(directive.Parameters) == 1 {
			location.Match = directive.Parameters[0]
			return location, nil
		} else if len(directive.Parameters) == 2 {
			location.Modifier = directive.Parameters[0]
			location.Match = directive.Parameters[1]
			return location, nil
		}

		return location, nil
	}

	return nil, errors.New("location must have block")

}

func (l *Location) GetParameters() []string {
	return []string{l.Modifier, l.Match}
}

func (l *Location) GetName() string {
	return l.Name
}

func (l *Location) GetDirectives() []statement.IDirective {
	directives := make([]statement.IDirective, 0)
	directives = append(directives, l.Directives...)

	return directives
}

func (l *Location) FindDirectives(directiveName string) []statement.IDirective {
	directives := make([]statement.IDirective, 0)

	for _, directive := range l.GetDirectives() {
		if directive.GetName() == directiveName {
			directives = append(directives, directive)
		}
	}

	return directives
}

func (l *Location) GetBlock() statement.IBlock {
	return l
}
