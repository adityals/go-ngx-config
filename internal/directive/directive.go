package directive

import "github.com/adityalstkp/go-ngx-config/internal/statement"

type Directive struct {
	Block      statement.IBlock
	Name       string
	Parameters []string
}

func (d *Directive) GetName() string {
	return d.Name
}

func (d *Directive) GetParameters() []string {
	return d.Parameters
}

func (d *Directive) GetBlock() statement.IBlock {
	return d.Block
}
