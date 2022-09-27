package main

import (
	"github.com/adityalstkp/go-ngx-config/pkg/cli"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "go-ngx-config COMMAND [ARG...]",
		Version: cli.VERSION,
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
