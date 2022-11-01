package main

import (
	"errors"
	"time"

	"github.com/adityals/go-ngx-config/internal/matcher"
	"github.com/adityals/go-ngx-config/pkg/parser"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func RunNgxLocationTester(cmd *cobra.Command, args []string) error {
	startTime := time.Now()

	filePath, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}

	targetUrl, err := cmd.Flags().GetString("url")
	if err != nil {
		return err
	}

	parserOpts := parser.NgxConfParserCliOptions{
		Filepath: filePath,
	}

	ast, err := parser.NewNgxConfParser(parserOpts)
	if err != nil {
		return err
	}

	match, err := matcher.NewLocationMatcher(ast, targetUrl)
	if err != nil {
		return err
	}

	if match == nil {
		return errors.New("match not found")
	}

	elapsed := time.Since(startTime)
	logrus.Info("Process time: ", elapsed)

	logrus.Info("[Match] Modifier: ", match.MatchModifer)
	logrus.Info("[Match] Path: ", match.MatchPath)

	return nil

}
