package main

import (
	"fmt"
	"strconv"
	"strings"
)

// a (chain of) dice rolls (lets say "dicechain") result in a string like
// "12345" (the first throw was a '1', the second one a '2' etc.).
//
// a diceware list essentially is a list of words, prefixed with a dice roll
// chain:
//
//	11111	a
//	11112	a&p
//	11113	a's
//	11114	aa
//	11115	aaa
//	11116	aaaa
//
// so, '11111' maps to index '0', '11116' maps to index '6' in the list.
//
// by using (real world) dices, a user creates a dicechain, look up the
// word in the list, repeat. the -electronic mode switches to the weaker
// (from a quality of randomness point of view) (pseudo-) random number
// generator of the operating system, picks the word of the used list
// based upon the random number, repeat.
//
// i removed the dicechain prefix of each list to slim down the size of
// each list.

// parse a dicechain into an "index".
func parseDiceChain(chain string) (int, error) {

	if len(chain) != 5 {
		return -1, fmt.Errorf("invalid dicechain %q: expected 6 chars", chain)
	}

	r := strings.Map(diceDown, chain)
	n, err := strconv.ParseInt(r, 6, 32)

	return int(n), err
}

func indexToDiceChain(index int) string {

	r := strconv.FormatInt(int64(index), 6)
	r = fmt.Sprintf("%05s", r)
	return strings.Map(diceUp, r)
}

func diceDown(r rune) rune {

	switch r {
	case '1':
		return '0'
	case '2':
		return '1'
	case '3':
		return '2'
	case '4':
		return '3'
	case '5':
		return '4'
	case '6':
		return '5'
	}

	panic("unhandled rune " + string(r))

	return 0
}

func diceUp(r rune) rune {

	switch r {
	case '0':
		return '1'
	case '1':
		return '2'
	case '2':
		return '3'
	case '3':
		return '4'
	case '4':
		return '5'
	case '5':
		return '6'
	}

	panic("unhandled rune " + string(r))

	return 0
}
