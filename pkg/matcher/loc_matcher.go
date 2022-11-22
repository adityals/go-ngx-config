package matcher

import (
	"errors"

	"github.com/adityals/go-ngx-config/internal/crossplane"
	"github.com/adityals/go-ngx-config/internal/matcher"
)

func NewLocationMatcher(filename string, targetUrl string, opts *crossplane.ParseOptions) (*matcher.LocationMatcher, error) {
	payload, err := crossplane.Parse(filename, opts)
	if err != nil {
		return nil, err
	}

	match, err := matcher.NewLocationMatcher(payload, targetUrl)
	if err != nil {
		return nil, err
	}

	if match == nil {
		return nil, errors.New("match is nil")
	}

	return match, nil
}

func NewLocationMatcherFromPayload(payload *crossplane.Payload, targetUrl string) (*matcher.LocationMatcher, error) {
	match, err := matcher.NewLocationMatcher(payload, targetUrl)
	if err != nil {
		return nil, err
	}

	if match == nil {
		return nil, errors.New("match is nil")
	}

	return match, nil
}
