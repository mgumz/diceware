package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type wordPrinter struct {
	horizontal bool
	verbose    bool
}

func (p *wordPrinter) Print(words []wordLine) {

	w := &tabwriter.Writer{}
	w.Init(os.Stdout, 0, 8, 2, ' ', tabwriter.AlignRight)
	defer w.Flush()

	if p.horizontal {
		for i := range words {
			fmt.Fprintf(w, "%s\t", words[i].word)
		}
		fmt.Fprintln(w)
		if p.verbose {
			for i := range words {
				fmt.Fprintf(w, "%s\t", indexToDiceChain(words[i].index))
			}
			fmt.Fprintln(w)
		}
	} else {
		for i := range words {
			fmt.Fprintf(w, "%s", words[i].word)
			if p.verbose {
				fmt.Fprintf(w, "\t%s", indexToDiceChain(words[i].index))
			}
			fmt.Fprintln(w)
		}
	}
}
