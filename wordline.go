package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

//
// ========================================================================
//

type wordLine struct {
	word  string
	index int
}

func (wl *wordLine) Set(line string, index int) {
	wl.word = line
	wl.index = index
}

//
// ========================================================================
//

var internalLists = make(map[string]string)

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
	return ioutil.NopCloser(r)
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
