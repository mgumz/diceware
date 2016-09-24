package main

import (
	"crypto/rand"
	"math/big"
)

func doRollOnList(rolls int, lines []string, printer *wordPrinter) {

	words := make([]wordLine, rolls)
	max := big.NewInt(int64(len(lines)))

	for i := 0; i < rolls; i += 1 {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic(err)
		}
		line := lines[int(n.Int64())]
		words[i].Set(line)
	}

	printer.Print(words)
}
