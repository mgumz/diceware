package main

import (
	"crypto/rand"
	"math/big"
)

func doRollOnList(rolls int, lines []string, extra bool, printer *wordPrinter) {

	words := make([]wordLine, rolls)
	max := len(lines)

	for i := 0; i < rolls; i += 1 {
		index := randN(max)
		line := lines[index]
		words[i].Set(line, index)
	}

	if extra {
		doExtra(words)
	}

	printer.Print(words)
}

// doExtra replace one char of one randomly selected word of words according
// to the diceware-extra rules. to cite diceware:
//
// For extra security without adding another word, insert one special
// character or digit chosen at random into your passphrase. Here is how to
// do this securely: Roll one die to choose a word in your passphrase, roll
// again to choose a letter in that word. Roll a third and fourth time to
// pick the added character from the following table:
//
//     Third Roll
//     1 2 3 4 5 6
// F 1 ~ ! # $ % ^
// o 2 & * ( ) - =
// u 3 + [ ] \ { }
// r 4 : ; " ' < >
// t 5 ? / 0 1 2 3
// h 6 4 5 6 7 8 9
//
func doExtra(words []wordLine) {

	const extra = "~!#$%^&*()-=+[]\\{}:;\"'<>?/0123456789"

	i := randN(len(words))
	word := []byte(words[i].word)
	word[randN(len(word))] = extra[randN(len(extra))]
	words[i].word = string(word)
}

func randN(max int) int {
	m := big.NewInt(int64(max))
	n, err := rand.Int(rand.Reader, m)
	if err != nil {
		panic(err)
	}
	return int(n.Int64())
}
