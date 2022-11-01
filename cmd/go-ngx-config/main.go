package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)

	rootCmd := NewRootCommand()
	parseCmd := NewParseCommand()
	locationTesterCmd := NewLocationTesterCommand()

	rootCmd.AddCommand(parseCmd)
	rootCmd.AddCommand(locationTesterCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
