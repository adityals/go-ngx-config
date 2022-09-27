package ast

import (
	"errors"

	"github.com/adityalstkp/go-ngx-config/internal/statement"
)

type Http struct {
	Servers    []*Server
	Name       string
	Directives []statement.IDirective
}

func NewHttp(directive statement.IDirective) (*Http, error) {
	if block := directive.GetBlock(); block != nil {
		http := &Http{
			Servers:    []*Server{},
			Directives: []statement.IDirective{},
			Name:       "http",
		}
		for _, directive := range block.GetDirectives() {
			if server, ok := directive.(*Server); ok {
				http.Servers = append(http.Servers, server)
				continue
			}
			http.Directives = append(http.Directives, directive)
		}

		return http, nil
	}

	return nil, errors.New("http directive must have a block")
}

func (h *Http) GetName() string {
	return h.Name
}

func (h *Http) GetParameters() []string {
	return []string{}
}

func (h *Http) GetDirectives() []statement.IDirective {
	directives := make([]statement.IDirective, 0)
	directives = append(directives, h.Directives...)

	for _, directive := range h.Servers {
		directives = append(directives, directive)
	}

	return directives
}

func (h *Http) FindDirectives(directiveName string) []statement.IDirective {
	directives := make([]statement.IDirective, 0)

	for _, directive := range h.GetDirectives() {
		if directive.GetName() == directiveName {
			directives = append(directives, directive)
		}
	}

	return directives
}

func (h *Http) GetBlock() statement.IBlock {
	return h
}
