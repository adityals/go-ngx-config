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

	logrus.Info("Test location match")

	filePath, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}

	parseInclude, err := cmd.Flags().GetBool("include")
	if err != nil {
		return err
	}

	targetUrl, err := cmd.Flags().GetString("url")
	if err != nil {
		return err
	}

	logrus.Info("Parsing include: ", parseInclude)

	parserOpts := parser.NgxConfParserCliOptions{
		Filepath:     filePath,
		ParseInclude: parseInclude,
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

	logrus.Info("[Match] Modifier: ", match.MatchModifer)
	logrus.Info("[Match] Path: ", match.MatchPath)

	logrus.Info("[Match] --- Directives Inside --- ")
	for _, d := range match.Directives {
		logrus.Info("[Match] Name: ", d.GetName())
		for _, param := range d.GetParameters() {
			logrus.Info("[Match] Parameters: ", param)
		}
	}
	logrus.Info("[Match] --- End of Directives Inside --- ")

	logrus.Info("Process time: ", elapsed)

	return nil

}
