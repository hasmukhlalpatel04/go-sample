package main

import (
	"fmt"
)

func main() {

	hc := HelpChecker{
		HelpText: `Usage: myapp [options]
Options:
  -h, --help    Show this help message
  -f <file>     Specify input file
  -v            Enable verbose mode`,
	}

	if hc.CheckArgs() {
		// Help was shown, exit or rerun logic
		return
	}

	// Normal program execution continues here
	fmt.Println("Running main program...")

	fp := FileProcessor{FilePath: "example.txt"}

	// Predicate: keep lines containing "Go"
	predicate := func(line string) bool {
		return len(line) >= 2 && (line == "Go" || line == "golang" || line == "GO")
	}

	// Find all matching lines
	allMatches, err := fp.FindAll(predicate)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("All matches:", allMatches)

	// Read first matching line
	firstMatch, err := fp.ReadFirst(predicate)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("First match:", firstMatch)

	// Append matches back to file
	if err := fp.AppendToFile(allMatches); err != nil {
		fmt.Println("Error appending:", err)
		return
	}
	fmt.Println("Lines appended successfully.")
}
