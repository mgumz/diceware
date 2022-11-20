package main

import "flag"

type cliFlags struct {
	printVersion   bool
	listLists      bool
	dumpList       bool
	rollElectronic bool
	nRolls         int

	wordFile string
	wordList string
	extra    bool

	printUnpadded   bool
	printHorizontal bool
	wordSep         string
	beVerbose       bool
}

func setupFlags(cli *cliFlags) {
	flag.BoolVar(&cli.printVersion, "version", false,
		"show version, combine with -version")
	flag.BoolVar(&cli.listLists, "lists", false,
		"list internal lists")
	flag.BoolVar(&cli.dumpList, "dump", false,
		"dump the content of a -list, combine with -horizontal, -verbose")
	flag.StringVar(&cli.wordList, "list", "diceware",
		"name of list to use, see -lists")
	flag.IntVar(&cli.nRolls, "rolls", 6,
		"number of rolls for -electronic")
	flag.BoolVar(&cli.rollElectronic, "electronic", false,
		"roll dice electronically, see diceware FAQ")
	flag.StringVar(&cli.wordFile, "file", "",
		"read word list from file")
	flag.BoolVar(&cli.beVerbose, "verbose", false,
		"be more verbose (print line number of used word)")
	flag.BoolVar(&cli.extra, "extra", false,
		"modify one word according to 'extra' rules (-electronic)")
	flag.BoolVar(&cli.printHorizontal, "horizontal", true,
		"list rolled dice horizontally (-electronic)")
	flag.BoolVar(&cli.printUnpadded, "unpadded-words", false,
		"no padding while showing words")
	flag.StringVar(&cli.wordSep, "word-separator", " ",
		"separator between -unpadded words")
}
