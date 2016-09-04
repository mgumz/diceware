package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func doRollOnList(rolls int, lines []string) {
	max := big.NewInt(int64(len(lines)))

	for i := 0; i < rolls; i += 1 {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic(err)
		}
		fmt.Println(lines[int(n.Int64())])
	}
}
