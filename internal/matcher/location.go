package matcher

import (
	"errors"
	"net/url"
	"regexp"
	"strings"

	"github.com/adityals/go-ngx-config/internal/crossplane"
)

type LocationMatcher struct {
	MatchPath    string
	MatchModifer string
	Directives   crossplane.Directive
}

type locationDirective struct {
	Modifier   string
	Path       string
	Directives crossplane.Directive
}

const (
	EXACT                   = "="
	REGEX                   = "~"
	REGEX_NO_CASE_SENSITIVE = "~*"
	PREFIX_PRIORITY         = "^~"
	PREFIX                  = ""
)

func getLocation(directive []crossplane.Directive, locationDirectives *[]crossplane.Directive) {
	for _, parsed := range directive {
		if parsed.Directive == "location" {
			*locationDirectives = append(*locationDirectives, parsed)
		}

		if parsed.Block != nil && parsed.Directive != "location" {
			getLocation(*parsed.Block, locationDirectives)
		}

	}
}

func NewLocationMatcher(conf *crossplane.Payload, targetPath string) (*LocationMatcher, error) {
	if conf == nil {
		return nil, errors.New("no config can be compute")
	}

	parsedUrl, err := url.Parse(targetPath)
	if err != nil {
		return nil, err
	}

	locations := make([]locationDirective, 0)
	locationDirectives := make([]crossplane.Directive, 0)

	for _, v := range conf.Config {
		getLocation(v.Parsed, &locationDirectives)
	}

	if len(locationDirectives) == 0 {
		return nil, errors.New("no location(s) found")
	}

	for _, directive := range locationDirectives {
		args := directive.Args
		if len(args) == 1 {
			path := args[0]
			locations = append(locations, locationDirective{
				Directives: directive,
				Modifier:   "",
				Path:       path,
			})
			continue
		}

		modifier := args[0]
		path := args[1]
		locations = append(locations, locationDirective{
			Directives: directive,
			Modifier:   modifier,
			Path:       path,
		})

	}

	match, err := locationTester(locations, parsedUrl.Path)
	if err != nil {
		return nil, err
	}

	if match == nil {
		return nil, errors.New("no match found")

	}

	return &LocationMatcher{
		MatchModifer: match.MatchModifer,
		MatchPath:    match.MatchPath,
		Directives:   match.Directives,
	}, nil
}

func locationTester(locationsTarget []locationDirective, targetPath string) (*LocationMatcher, error) {
	// handle exact
	for _, location := range locationsTarget {
		if location.Modifier != EXACT {
			continue
		}

		if location.Path == targetPath {
			return &LocationMatcher{
				MatchPath:    location.Path,
				MatchModifer: location.Modifier,
				Directives:   location.Directives,
			}, nil
		}
	}

	// handle prefix and prefix priority
	var bestMatch locationDirective
	bestLength := 0
	for _, location := range locationsTarget {
		if location.Modifier != PREFIX && location.Modifier != PREFIX_PRIORITY {
			continue
		}

		if strings.HasPrefix(targetPath, location.Path) {
			locationLength := len(location.Path)
			if locationLength > bestLength {
				bestMatch = location
				bestLength = locationLength
			}
		}
	}

	// do not go to regex if priority
	if bestMatch.Path != "" && bestMatch.Modifier == PREFIX_PRIORITY {
		return &LocationMatcher{
			MatchPath:    bestMatch.Path,
			MatchModifer: bestMatch.Modifier,
			Directives:   bestMatch.Directives,
		}, nil
	}

	// handle regex
	for _, location := range locationsTarget {
		if location.Modifier == REGEX || location.Modifier == REGEX_NO_CASE_SENSITIVE {
			locationRegex := location.Path
			if location.Modifier == REGEX_NO_CASE_SENSITIVE {
				locationRegex = "(?i)" + locationRegex
			}

			reg, err := regexp.Compile(locationRegex)
			if err != nil {
				return nil, err
			}

			match := reg.FindString(targetPath)
			if match != "" {
				return &LocationMatcher{
					MatchPath:    location.Path,
					MatchModifer: location.Modifier,
					Directives:   location.Directives,
				}, nil
			}
		}
	}

	// use longest match
	if bestMatch.Path != "" {
		return &LocationMatcher{
			MatchPath:    bestMatch.Path,
			MatchModifer: bestMatch.Modifier,
			Directives:   bestMatch.Directives,
		}, nil
	}

	return nil, nil
}
