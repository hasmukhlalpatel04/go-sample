package main

import (
	"fmt"
	"os"
)

// HelpChecker encapsulates help flag detection
type HelpChecker struct {
	HelpText string
}

// CheckArgs inspects command-line arguments for -h or --help.
// If found, it prints the help text and returns true.
func (hc *HelpChecker) CheckArgs() bool {
	for _, arg := range os.Args[1:] {
		if arg == "-h" || arg == "--help" {
			fmt.Println(hc.HelpText)
			return true
		}
	}
	return false
}
