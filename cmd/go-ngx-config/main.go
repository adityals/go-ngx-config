package main

import (
	"fmt"
	"os"
)

func main() {
	rootCmd := NewRootCommand()
	parseCmd := NewParseCommand()

	rootCmd.AddCommand(parseCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
