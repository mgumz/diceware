package main

import (
	"bufio"
	"compress/gzip"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

var internalLists = make(map[string]string)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n%s [opts] [roll1] [roll2] [roll3] [...]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
	}

	list := flag.String("list", "diceware", "name of list to use, see -lists")
	listLists := flag.Bool("lists", false, "list internal lists")
	dumpList := flag.Bool("dump", false, "dump the content of a -list")
	rolls := flag.Int("rolls", 6, "number of rolls for -electronic")
	electronic := flag.Bool("electronic", false, "roll dice electronically (see diceware FAQ)")
	listFile := flag.String("file", "", "read word list from file")
	horizontal := flag.Bool("horizontal", true, "list rolled dice horizontally (-electronic)")
	verbose := flag.Bool("verbose", false, "be more verbose (print line number of used word)")
	extra := flag.Bool("extra", false, "modify one word according to 'extra' rules (-electronic)")

	flag.Parse()

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
		doDumpList(*list)
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

	doLookupRolls(flag.Args(), lines, printer)
}

func linesFromInternalList(list string) []string {
	r := internalListReader(list)
	if r == nil {
		return nil
	}
	defer r.Close()
	return readerToLines(r)
}

func readerToLines(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func internalListReader(name string) io.ReadCloser {
	list, exists := internalLists[name]
	if !exists {
		return nil
	}

	return listReaderFromString(list)
}

func listReaderFromString(list string) io.ReadCloser {
	r := strings.NewReader(strings.TrimSpace(list))
	b64 := base64.NewDecoder(base64.StdEncoding, r)
	gz, err := gzip.NewReader(b64)
	if err != nil {
		panic(err)
	}
	return gz
}

func internalListNames() []string {
	names := make([]string, len(internalLists))
	i := 0
	for name, _ := range internalLists {
		names[i] = name
		i += 1
	}
	sort.Strings(names)
	return names
}
