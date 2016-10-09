package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// dumps a given word list, identified by 'name', to stdout
// 'horizontal' - prints 6 words (+ the index) on each line
// 'verbose' - prints also the available meta data befor all words
func doDumpList(name string, horizontal, verbose bool) {
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

	if horizontal {
		for i, word := range linesFromInternalList(name) {
			fmt.Fprint(w, list.Index(i)+" "+word)
			if (i+1)%6 == 0 {
				fmt.Fprintln(w)
			} else {
				fmt.Fprint(w, "\t")
			}
		}

	} else {
		for i, word := range linesFromInternalList(name) {
			fmt.Fprintln(w, list.Index(i), "\t", word)
		}
	}

}
