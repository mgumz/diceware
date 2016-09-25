package main

import "fmt"

func doLookupRolls(rolls, lines []string, printer *wordPrinter) error {

	words := make([]wordLine, len(rolls))

	for i, chain := range rolls {
		index, err := parseDiceChain(chain)
		if err != nil {
			return fmt.Errorf("%q not a valid dicechain\n", chain)
		}
		if index < 0 || index >= len(lines) {
			return fmt.Errorf("%q not in the list\n", chain)
		}

		word := lines[index]
		words[i].Set(word, index)
	}

	printer.Print(words)

	return nil
}
