package matcher

import (
	"errors"

	"github.com/adityals/go-ngx-config/internal/matcher"
	"github.com/adityals/go-ngx-config/internal/parser"
)

func NewLocationMatcher(locMatcherOpts LocationMatcherOptions) (*matcher.LocationMatcher, error) {
	parser, err := parser.NewParser(locMatcherOpts.Filepath)
	if err != nil {
		return nil, err
	}

	parsedConf := parser.Parse()
	if parsedConf == nil {
		return nil, errors.New("cannot be parsed")
	}

	match, err := matcher.NewLocationMatcher(parsedConf, locMatcherOpts.UrlTarget)
	if err != nil {
		return nil, err
	}

	if match == nil {
		return nil, errors.New("match is nil")
	}

	return match, nil
}
