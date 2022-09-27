package main

import (
	"encoding/json"

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
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath, err := cmd.Flags().GetString("file")
			if err != nil {
				return err
			}

			ast, err := cli.NewNgxConfParser(filePath)
			if err != nil {
				return err
			}

			ast_json, err := json.MarshalIndent(ast, "", "  ")
			if err != nil {
				return err
			}

			println(string(ast_json))
			return nil

		},
	}

	parseCmd.Flags().StringP("file", "f", "", "nginx.conf file location")

	return parseCmd
}
