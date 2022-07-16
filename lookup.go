package main

import "fmt"

// for each dice roll in a list of rolls, lookup the appropriate
// word in 'words' and feeds the selected words to 'printer'
func doLookupRolls(rolls, words []string, printer *wordPrinter) error {

	selected := make([]wordLine, len(rolls))

	for i, chain := range rolls {
		index, err := diceChain(chain).toIndex()
		if err != nil {
			return fmt.Errorf("%q is NOT a valid dicechain", chain)
		}
		if index < 0 || index >= len(words) {
			return fmt.Errorf("%q is NOT in the list", chain)
		}

		word := words[index]
		selected[i].Set(word, index)
	}

	printer.Print(selected)

	return nil
}
