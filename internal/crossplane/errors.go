package crossplane

import (
	"fmt"
)

type ParseError struct {
	what string
	file *string
	line *int
}

func (e ParseError) Error() string {
	if e.line != nil && e.file != nil {
		return fmt.Sprintf("%s in %s:%d", e.what, *e.file, *e.line)
	}

	if e.line != nil {
		return fmt.Sprintf("%s in %d", e.what, *e.line)
	}

	return e.what
}
