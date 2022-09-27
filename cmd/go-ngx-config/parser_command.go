package main

import (
	"encoding/json"

	"github.com/adityalstkp/go-ngx-config/pkg/cli"
	"github.com/spf13/cobra"
)

func RunParseNgx(cmd *cobra.Command, args []string) error {
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

}
