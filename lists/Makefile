#
# http://world.std.com/~reinhold/diceware.html#Diceware in Other Languages|outline
#

ALL=reinhold.gz.b64 \
	beale.gz.b64 \
	german.gz.b64 \
	russian.gz.b64 \
	eff2016.gz.b64

GZ_B64=gzip -9 -c | ENC=64 WRAP=78 r2t


all: $(ALL)
clean:
	rm -f $(ALL)


reinhold.gz.b64: diceware.wordlist.asc
	sed -n -e '3,7778p' $< | $(GZ_B64) > $@

beale.gz.b64: beale.wordlist.asc
	sed -n -e '3,7778p' $< | $(GZ_B64) > $@

german.gz.b64: diceware_german.txt
	cat $< | $(GZ_B64) > $@

russian.gz.b64: diceware.ru.txt
	cat $< | $(GZ_B64) > $@

eff2016.gz.b64: eff_large_wordlist.txt
	cat $< | $(GZ_B64) > $@
