package ast

import (
	"github.com/adityalstkp/go-ngx-config/internal/statement"
)

type Block struct {
	Directives []statement.IDirective
}

func (b *Block) GetDirectives() []statement.IDirective {
	return b.Directives
}

func (b *Block) FindDirectives(directiveName string) []statement.IDirective {
	directives := make([]statement.IDirective, 0)

	for _, directive := range b.GetDirectives() {
		if directive.GetName() == directiveName {
			directives = append(directives, directive)
		}

		if include, ok := directive.(*Include); ok {
			for _, c := range include.Configs {
				directives = append(directives, c.FindDirectives(directiveName)...)
			}
		}

		if directive.GetBlock() != nil {
			directives = append(directives, directive.GetBlock().FindDirectives(directiveName)...)
		}
	}

	return directives
}
