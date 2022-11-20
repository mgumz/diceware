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

	cli := cliFlags{}
	setupFlags(&cli)

	flag.Parse()

	if cli.printVersion {
		printVersion(cli.beVerbose)
		return
	}

	if cli.listLists {
		for _, name := range internalListNames() {
			fmt.Println(name)
		}
		return
	}

	if cli.dumpList {
		if cli.wordList == "" {
			fmt.Fprintf(os.Stderr, "error: missing -list\n")
			return
		}
		doDumpList(cli.wordList, cli.printHorizontal, cli.beVerbose)
		return
	}

	lines, err := getWordLines(cli.wordFile, cli.wordList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(2)
	}

	printer := &wordPrinter{
		os.Stdout,
		cli.printUnpadded, cli.wordSep,
		!cli.printHorizontal,
		cli.beVerbose,
	}

	if cli.rollElectronic {
		doRollOnList(cli.nRolls, lines, cli.extra, printer)
		return
	}

	if len(flag.Args()) == 0 {
		fmt.Println("no rolls given, no lookup possible.")
		flag.Usage()
		return
	}

	if err := doLookupRolls(flag.Args(), lines, printer); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func getWordLines(filename string, listName string) ([]string, error) {
	if filename != "" {
		return getWordLinesFromFile(filename)
	}
	return getWordLinesFromInternal(listName)
}

func getWordLinesFromFile(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return []string{}, fmt.Errorf("%q: %v", name, err)
	}
	defer file.Close()
	return readerToLines(file), nil
}

func getWordLinesFromInternal(name string) ([]string, error) {
	lines := linesFromInternalList(name)
	if lines == nil {
		return []string{}, fmt.Errorf("list %q does not exist", name)
	}
	return lines, nil
}
