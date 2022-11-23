package main

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/adityals/go-ngx-config/internal/crossplane"
	"github.com/adityals/go-ngx-config/pkg/parser"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func RunParseNgx(cmd *cobra.Command, args []string) error {
	startTime := time.Now()

	logrus.Info("Parsing nginx config")

	filePath, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}

	singleFile, err := cmd.Flags().GetBool("single")
	if err != nil {
		return err
	}

	outputFilePath, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	logrus.Info("Single File: ", singleFile)

	ast, err := parser.NewNgxConfParser(filePath, &crossplane.ParseOptions{
		SingleFile:     singleFile,
		CombineConfigs: true,
	})
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
