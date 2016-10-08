package main

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
