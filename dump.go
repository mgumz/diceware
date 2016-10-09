package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func doDumpList(name string, verbose bool) {
	list, exists := internalLists[name]
	if !exists {
		fmt.Fprintf(os.Stderr, "error: list %q does not exist\n", name)
		return
	}

	w := &tabwriter.Writer{}
	w.Init(os.Stdout, 0, 8, 2, ' ', 0)
	defer w.Flush()

	if verbose {

		fmt.Fprintln(w, "Name:", "\t", list.Name)
		fmt.Fprintln(w, "Author:", "\t", list.Author)
		fmt.Fprintln(w, "Origin:", "\t", list.Origin)
		fmt.Fprintln(w, "License:", "\t", list.License)

		fmt.Fprintln(w, "---")
	}

	for i, word := range linesFromInternalList(name) {
		fmt.Fprintln(w, list.Index(i), "\t", word)
	}

}
