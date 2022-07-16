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
// '11111' maps to index '0', '11116' maps to index '6' in the list, etc.
//
// since dices start with '1' while an index in golang starts with '0',
// we have to translate the dicechain into a proper base-6 number to get
// the index in the list of words. see diceDown() and diceUp() for the
// respective transformations.
//
// by using (real world) dices, a user creates a dicechain, looks up the
// word in the list, repeats. the -electronic mode switches to the weaker
// (from a quality of randomness point of view) (pseudo-) random number
// generator of the operating system, picks the word of the used list
// based upon the random number, repeat.
//
// i removed the dicechain prefix of each list to slim down the size of
// each list.
type diceChain string

func (dc diceChain) String() string { return string(dc) }

// transform a dicechain into an "index".
func (chain diceChain) toIndex() (int, error) {

	if len(chain) != 5 {
		return -1, fmt.Errorf("invalid dicechain %q: expected 5 chars per dicechain", chain)
	}

	r := strings.Map(diceDown, chain.String())
	if len(r) != 5 {
		return -1, fmt.Errorf("invalid dicechain %q: contains an invalid char", chain)
	}

	n, err := strconv.ParseInt(r, 6, 32)

	return int(n), err
}

func (dc diceChain) FromIndex(index int) diceChain {
	r := strconv.FormatInt(int64(index), 6)
	r = fmt.Sprintf("%05s", r)
	dc = diceChain(strings.Map(diceUp, r))
    return dc
}

// transform '12345' to '01234' - all numbers down by 1
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
	default:
		return -1
	}
}

// inverse of diceDown
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
	default:
		return -1
	}
}
