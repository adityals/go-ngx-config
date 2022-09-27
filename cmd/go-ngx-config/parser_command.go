package main

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/adityalstkp/go-ngx-config/pkg/cli"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func RunParseNgx(cmd *cobra.Command, args []string) error {
	startTime := time.Now()

	filePath, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}

	outputFilePath, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	cliOpts := cli.NgxConfParserCliOptions{
		Filepath: filePath,
	}

	ast, err := cli.NewNgxConfParser(cliOpts)
	if err != nil {
		return err
	}

	ast_json, err := json.MarshalIndent(ast, "", "  ")
	if err != nil {
		return err
	}

	if outputFilePath != "" {
		if _, err := os.Stat(outputFilePath); errors.Is(err, os.ErrNotExist) {
			err := os.MkdirAll(outputFilePath, os.ModePerm)
			if err != nil {
				return err
			}
		}

		dumpAstJsonFile := outputFilePath + "/dump.json"
		f, err := os.Create(dumpAstJsonFile)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.Write(ast_json)
		if err != nil {
			return err
		}
	} else {
		println(string(ast_json))
	}

	elapsed := time.Since(startTime)
	logrus.Info("Process time: ", elapsed)

	return nil

}
