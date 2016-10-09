package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n%s [opts] [roll1] [roll2] [roll3] [...]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
	}

	version := flag.Bool("version", false, "show version, combine with -version")
	list := flag.String("list", "diceware", "name of list to use, see -lists")
	listLists := flag.Bool("lists", false, "list internal lists")
	dumpList := flag.Bool("dump", false, "dump the content of a -list, combine with -horizontal,-verbose")
	rolls := flag.Int("rolls", 6, "number of rolls for -electronic")
	electronic := flag.Bool("electronic", false, "roll dice electronically, see diceware FAQ")
	listFile := flag.String("file", "", "read word list from file")
	horizontal := flag.Bool("horizontal", true, "list rolled dice horizontally (-electronic)")
	verbose := flag.Bool("verbose", false, "be more verbose (print line number of used word)")
	extra := flag.Bool("extra", false, "modify one word according to 'extra' rules (-electronic)")

	flag.Parse()

	if *version {
		printVersion(*verbose)
		return
	}

	if *listLists {
		for _, name := range internalListNames() {
			fmt.Println(name)
		}
		return
	}

	if *dumpList {
		if *list == "" {
			fmt.Fprintf(os.Stderr, "error: missing -list\n")
			return
		}
		doDumpList(*list, *horizontal, *verbose)
		return
	}

	var lines []string
	if *listFile != "" {
		file, err := os.Open(*listFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %q: %v\n", err)
			os.Exit(2)
		}
		defer file.Close()
		lines = readerToLines(file)
	} else {
		if lines = linesFromInternalList(*list); lines == nil {
			fmt.Fprintf(os.Stderr, "error: list %q does not exist\n", *list)
			os.Exit(1)
		}
	}

	printer := &wordPrinter{*horizontal, *verbose}

	if *electronic {
		doRollOnList(*rolls, lines, *extra, printer)
		return
	}

	if len(flag.Args()) == 0 {
		fmt.Println("no rolls given, no lookup possible.\n")
		flag.Usage()
		return
	}

	if err := doLookupRolls(flag.Args(), lines, printer); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
}
