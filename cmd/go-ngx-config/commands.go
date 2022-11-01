package main

import (
	"github.com/adityals/go-ngx-config/pkg/parser"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "go-ngx-config COMMAND [ARG...]",
		Version: parser.VERSION,
		Short:   "A nginx config build with go",
	}

	return rootCmd
}

func NewParseCommand() *cobra.Command {
	parseCmd := &cobra.Command{
		Use:   "parse",
		Short: "A nginx config parser",
		RunE:  RunParseNgx,
	}

	parseCmd.Flags().StringP("file", "f", "", "nginx.conf file location")
	parseCmd.Flags().StringP("output", "o", "", "output file location")

	return parseCmd
}

func NewLocationTesterCommand() *cobra.Command {
	testCmd := &cobra.Command{
		Use:   "lt",
		Short: "A nginx location tester",
		RunE:  RunNgxLocationTester,
	}

	testCmd.Flags().StringP("file", "f", "", "nginx.conf file location")
	testCmd.Flags().StringP("url", "u", "", "target url")

	return testCmd
}
