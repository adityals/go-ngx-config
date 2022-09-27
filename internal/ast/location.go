package ast

import (
	"errors"

	"github.com/adityals/go-ngx-config/internal/directive"
)

type Location struct {
	*directive.Directive
	Name     string
	Modifier string
	Match    string
}

func NewLocation(directive *directive.Directive) (*Location, error) {
	location := &Location{
		Name: "location",
	}

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
