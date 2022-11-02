package matcher

import (
	"errors"
	"net/url"
	"regexp"
	"strings"

	"github.com/adityals/go-ngx-config/internal/ast"
	"github.com/adityals/go-ngx-config/internal/statement"
)

type LocationMatcher struct {
	MatchPath    string
	MatchModifer string
	Directives   []statement.IDirective
}

const (
	EXACT                   = "="
	REGEX                   = "~"
	REGEX_NO_CASE_SENSITIVE = "~*"
	PREFIX_PRIORITY         = "^~"
	PREFIX                  = ""
)

func NewLocationMatcher(conf *ast.Config, targetPath string) (*LocationMatcher, error) {
	if conf == nil {
		return nil, errors.New("no config can be compute")
	}

	parsedUrl, err := url.Parse(targetPath)
	if err != nil {
		return nil, err
	}

	locations := make([]ast.Location, 0)

	locationDirectives := conf.FindDirectives("location")
	if len(locationDirectives) == 0 {
		return nil, errors.New("no location(s) found")
	}

	for _, directive := range locationDirectives {
		name := directive.GetName()
		parameters := directive.GetParameters()
		modifier := parameters[0]
		match := parameters[1]
		directives := directive.GetBlock().GetDirectives()

		locations = append(locations, ast.Location{
			Name:       name,
			Modifier:   modifier,
			Match:      match,
			Directives: directives,
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

func locationTester(locationsTarget []ast.Location, targetPath string) (*LocationMatcher, error) {
	// handle exact
	for _, location := range locationsTarget {
		if location.Modifier != EXACT {
			continue
		}

		if location.Match == targetPath {
			return &LocationMatcher{
				MatchPath:    location.Match,
				MatchModifer: location.Modifier,
				Directives:   location.Directives,
			}, nil
		}
	}

	// handle prefix and prefix priority
	var bestMatch ast.Location
	bestLength := 0
	for _, location := range locationsTarget {
		if location.Modifier != PREFIX && location.Modifier != PREFIX_PRIORITY {
			continue
		}

		if strings.HasPrefix(targetPath, location.Match) {
			locationLength := len(location.Match)
			if locationLength > bestLength {
				bestMatch = location
				bestLength = locationLength
			}
		}
	}

	// do not go to regex if priority
	if bestMatch.Match != "" && bestMatch.Modifier == PREFIX_PRIORITY {
		return &LocationMatcher{
			MatchPath:    bestMatch.Match,
			MatchModifer: bestMatch.Modifier,
			Directives:   bestMatch.Directives,
		}, nil
	}

	// handle regex
	for _, location := range locationsTarget {
		if location.Modifier == REGEX || location.Modifier == REGEX_NO_CASE_SENSITIVE {
			locationRegex := location.Match
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
					MatchPath:    location.Match,
					MatchModifer: location.Modifier,
					Directives:   location.Directives,
				}, nil
			}
		}
	}

	// use longest match
	if bestMatch.Match != "" {
		return &LocationMatcher{
			MatchPath:    bestMatch.Match,
			MatchModifer: bestMatch.Modifier,
			Directives:   bestMatch.Directives,
		}, nil
	}

	return nil, nil
}
