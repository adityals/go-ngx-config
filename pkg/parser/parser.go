package parser

import "github.com/adityals/go-ngx-config/internal/crossplane"

func NewNgxConfParser(filename string, opts *crossplane.ParseOptions) (*crossplane.Payload, error) {
	payload, err := crossplane.Parse(filename, opts)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func NewNgxConfStringParser(conf string, opts *crossplane.ParseOptions) (*crossplane.Payload, error) {
	payload, err := crossplane.ParseString(conf, opts)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
