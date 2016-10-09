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
	polish.stripped \
	dutch.stripped


all: $(ALL)
clean:
	rm -f $(ALL)

reinhold.stripped: diceware.txt
	sed -n -e '3,7778 s/^[0-9]*[ 	]//p' $< > $@

beale.stripped: diceware_beale.txt
	sed -n -e '3,7778 s/^[0-9]*[ 	]//p' $< > $@

japanese.stripped: diceware_jp.txt
	sed -n -e '8,7783 s/^[0-9]*[ 	]//p' $< > $@

italy.stripped: diceware_it.txt
	sed -n -e '18,7793 s/^[0-9]*[ 	]//p' $< > $@

dutch.stripped: diceware_nl.txt
	iconv -f latin1 -t utf8 $< | sed -n -e 's/^[0-9]*[ 	]//p' > $@

german.stripped: diceware_de.txt
	sed -n -e 's/^[0-9]*[ 	]//p' $< > $@

svedish.stripped: diceware_sv.txt
	sed -n -e 's/^[0-9]*[ 	]//p' $< > $@

turkish.stripped: diceware_tr.txt
	sed -n -e 's/^[0-9]*[ 	]//p' $< > $@

portuguese.stripped: diceware_pt.txt
	sed -n -e '4,7779 s/^[0-9]*[ 	]*//p' $< > $@

polish.stripped: dicelist_pl.txt
	sed -n -e '12,7787 s/^[0-9]*[ 	]*//p' $< > $@

russian.stripped: diceware_ru.txt
	sed -n -e 's/^[0-9]*[ 	]//p' $< | iconv -t latin1 | iconv -f cp1251 -t utf-8 > $@

eff2016.stripped: diceware_eff2016.txt
	sed -n -e 's/^[0-9]*[ 	]//p' $< > $@

