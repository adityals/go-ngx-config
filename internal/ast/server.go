package ast

import (
	"errors"

	"github.com/adityalstkp/go-ngx-config/internal/statement"
)

type Server struct {
	Block statement.IBlock
}

func NewServer(directive statement.IDirective) (*Server, error) {
	if block := directive.GetBlock(); block != nil {
		return &Server{
			Block: block,
		}, nil
	}

	return nil, errors.New("server must have a block")
}

func (s *Server) GetName() string {
	return "server"
}

func (s *Server) GetParameters() []string {
	return []string{}
}

func (s *Server) GetBlock() statement.IBlock {
	return s.Block
}
