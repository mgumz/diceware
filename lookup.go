package main

import (
	"fmt"
	"os"
	"strings"
)

func doLookupRolls(rolls, lines []string) {

	for _, roll := range rolls {
		line := lookup(roll, lines)
		if line != "" {
			fmt.Println(line)
		} else {
			fmt.Fprintf(os.Stderr, "error: %q not found in list", roll)
		}
	}
}

func lookup(roll string, lines []string) string {
	for _, line := range lines {
		if strings.HasPrefix(line, roll) {
			return line
		}
	}
	return ""
}
