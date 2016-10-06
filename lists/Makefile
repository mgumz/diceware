#
# http://world.std.com/~reinhold/diceware.html#Diceware in Other Languages|outline
#

ALL=reinhold.stripped \
	beale.stripped \
	german.stripped \
	japanese.stripped \
	italy.stripped \
	russian.stripped \
	eff2016.stripped

all: $(ALL)
clean:
	rm -f $(ALL)

reinhold.stripped: diceware.wordlist.asc
	sed -n -e '3,7778 s/^[0-9]*[ 	]//p' $< > $@

beale.stripped: beale.wordlist.asc
	sed -n -e '3,7778 s/^[0-9]*[ 	]//p' $< > $@

japanese.stripped: diceware_jp.txt
	sed -n -e '8,7783 s/^[0-9]*[ 	]//p' $< > $@

italy.stripped: word_list_diceware_it-IT-2.txt
	sed -n -e '18,7793 s/^[0-9]*[ 	]//p' $< > $@

german.stripped: diceware_german.txt
	sed -n -e 's/^[0-9]*[ 	]//p' $< > $@

russian.stripped: diceware.ru.txt
	sed -n -e 's/^[0-9]*[ 	]//p' $< > $@

eff2016.stripped: eff_large_wordlist.txt
	sed -n -e 's/^[0-9]*[ 	]//p' $< > $@
