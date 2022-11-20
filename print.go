package main

import (
	"fmt"
	"io"
	"text/tabwriter"
)

type wordPrinter struct {
	w        io.Writer
	simple   bool
	sep      string
	vertical bool
	verbose  bool
}

func (p *wordPrinter) Print(words []wordLine) {

	if p.simple {
		printSimple(p.w, words, p.sep)
		return
	}

	tw := &tabwriter.Writer{}
	tw.Init(p.w, 0, 8, 2, ' ', tabwriter.AlignRight)
	defer tw.Flush()

	if !p.vertical {
		for i := range words {
			fmt.Fprintf(tw, "%s\t", words[i].word)
		}
		fmt.Fprintln(tw)
		if p.verbose {
			for i := range words {
				fmt.Fprintf(tw, "%s\t", indexToDiceChain(words[i].index))
			}
			fmt.Fprintln(tw)
		}
	} else {
		for i := range words {
			fmt.Fprintf(tw, "%s", words[i].word)
			if p.verbose {
				fmt.Fprintf(tw, "\t%s", indexToDiceChain(words[i].index))
			}
			fmt.Fprintln(tw)
		}
	}
}

func printSimple(w io.Writer, words []wordLine, sep string) {

	for i := range words {
		fmt.Fprint(w, words[i].word)
		if i < len(words)-1 {
			fmt.Fprint(w, sep)
		} else {
			fmt.Fprintln(w)
		}
	}

}
