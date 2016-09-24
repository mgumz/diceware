package main

import (
	"fmt"
	"os"
	"strings"
)

func doLookupRolls(rolls, lines []string, printer *wordPrinter) {

	words := make([]wordLine, len(rolls))

	for i, roll := range rolls {
		line := lookup(roll, lines)
		if line != "" {
			words[i].Set(line)
		} else {
			fmt.Fprintf(os.Stderr, "error: %q not found in list\n", roll)
			return
		}
	}

	printer.Print(words)
}

func lookup(roll string, lines []string) string {
	for _, line := range lines {
		if strings.HasPrefix(line, roll) {
			return line
		}
	}
	return ""
}
