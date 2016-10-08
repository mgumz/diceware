package main

import "fmt"

func doLookupRolls(rolls, words []string, printer *wordPrinter) error {

	selected := make([]wordLine, len(rolls))

	for i, chain := range rolls {
		index, err := parseDiceChain(chain)
		if err != nil {
			return fmt.Errorf("%q not a valid dicechain\n", chain)
		}
		if index < 0 || index >= len(words) {
			return fmt.Errorf("%q not in the list\n", chain)
		}

		word := words[index]
		selected[i].Set(word, index)
	}

	printer.Print(selected)

	return nil
}
