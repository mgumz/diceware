#
# http://world.std.com/~reinhold/diceware.html#Diceware in Other Languages|outline
#

ALL=reinhold.stripped \
	beale.stripped \
	german.stripped \
	japanese.stripped \
	italy.stripped \
	russian.stripped \
	eff2016.stripped \
	svedish.stripped \
	turkish.stripped \
	portuguese.stripped \
	polish.stripped


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

svedish.stripped: diceware-sv.txt
	sed -n -e 's/^[0-9]*[ 	]//p' $< > $@

turkish.stripped: diceware_tr.txt
	sed -n -e 's/^[0-9]*[ 	]//p' $< > $@

portuguese.stripped: diceware.wordlist.pt.txt
	sed -n -e '4,7779 s/^[0-9]*[ 	]*//p' $< > $@

polish.stripped: dicelist-pl.txt
	sed -n -e '12,7787 s/^[0-9]*[ 	]*//p' $< > $@

russian.stripped: diceware.ru.txt
	sed -n -e 's/^[0-9]*[ 	]//p' $< | iconv -t latin1 | iconv -f cp1251 -t utf-8 > $@

eff2016.stripped: eff_large_wordlist.txt
	sed -n -e 's/^[0-9]*[ 	]//p' $< > $@

