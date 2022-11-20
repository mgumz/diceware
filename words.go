package main

import (
	"bufio"
	"io"
	"sort"
	"strings"
)

type WordList struct {
	Name    string
	Author  string
	License string
	Origin  string
	Index   IndexStringer
	Words   string
}

var internalLists = map[string]*WordList{}

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

	return listReaderFromString(list.Words)
}

func listReaderFromString(list string) io.ReadCloser {
	r := strings.NewReader(strings.TrimSpace(list))
	return io.NopCloser(r)
}

func internalListNames() []string {
	names := make([]string, len(internalLists))
	i := 0
	for name := range internalLists {
		names[i] = name
		i++
	}
	sort.Strings(names)
	return names
}
