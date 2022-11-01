package main

import (
	"github.com/adityals/go-ngx-config/internal/matcher"
	"github.com/adityals/go-ngx-config/pkg/parser"
)

func main() {
	conf, err := parser.NewStringNgxConfParser(`
			location ~ /test-mypath {
				return 200;
			}

			location = /test-match {
				return 200;
			}

			location / {
				return 200;
			}
	`)

	if err != nil {
		panic(err)
	}

	matcher, err := matcher.NewLocationMatcher(conf, "/test-match")
	if err != nil {
		panic(err)
	}

	println(matcher.MatchModifer, matcher.MatchPath)
}
