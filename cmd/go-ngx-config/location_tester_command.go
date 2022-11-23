package main

import (
	"errors"
	"time"

	"github.com/adityals/go-ngx-config/internal/crossplane"
	"github.com/adityals/go-ngx-config/pkg/matcher"
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

	singleFile, err := cmd.Flags().GetBool("single")
	if err != nil {
		return err
	}

	targetUrl, err := cmd.Flags().GetString("url")
	if err != nil {
		return err
	}

	logrus.Info("Single File: ", singleFile)

	match, err := matcher.NewLocationMatcher(filePath, targetUrl, &crossplane.ParseOptions{
		SingleFile:     singleFile,
		CombineConfigs: true,
	})
	if err != nil {
		return err
	}

	if match == nil {
		return errors.New("match not found")
	}

	elapsed := time.Since(startTime)

	logrus.Info("[Match] Modifier: ", match.MatchModifer)
	logrus.Info("[Match] Path: ", match.MatchPath)

	logrus.Info("[Match] --- Directives Inside Block --- ")
	for _, d := range *match.Directives.Block {
		logrus.Info("[Match] Name: ", d.Directive)
		for i, param := range d.Args {
			logrus.Infof("[Match] Args[%d]: %s", i, param)
		}
	}
	logrus.Info("[Match] --- End of Directives Inside Block --- ")

	logrus.Info("Process time: ", elapsed)

	return nil

}
