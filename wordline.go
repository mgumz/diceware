package main

import "strings"

type wordLine struct {
	word string
	line string
}

func (wl *wordLine) Set(line string) {
	field := strings.Fields(line)
	wl.word, wl.line = field[1], field[0]
}
