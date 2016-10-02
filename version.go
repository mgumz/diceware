package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

var (
	Version   = "0.1"
	GitHash   = ""
	BuildDate = ""
)

func printVersion(verbose bool) {

	tw := tabwriter.NewWriter(os.Stdout, 1, 2, 2, ' ', 0)

	fmt.Fprintln(tw, "diceware:\t"+Version)
	if verbose {
		if GitHash != "" {
			fmt.Fprintln(tw, "git:\t"+GitHash)
		}
		if BuildDate != "" {
			fmt.Fprintln(tw, "build-date:\t"+BuildDate)
		}
	}

	tw.Flush()
}
