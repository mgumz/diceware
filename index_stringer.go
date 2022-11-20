package main

import (
	"strconv"
)

// IndexStringer takes an "index" (a number) and transforms
// it into a string
type IndexStringer func(int) string

func indexToDiceChain(index int) string {
	dc := diceChain("").FromIndex(index)
	return dc.String()
}

func indexToPlainNumber(index int) string {
	return strconv.Itoa(index)
}
