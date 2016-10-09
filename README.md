# *diceware*

**diceware** creates passphrases along the guidelines posted on the [diceware]
page and [xkcd]. In short, **diceware** makes use of physical dice to select
words out of word lists. a word list consists of 7776 words, each word yields
~12.9 bits of entropy. Thus, a passphrase made out of 8 words gives a whopping
103.2 bits of entropy + the advantage of having something easy to remember.

**diceware** can be used to lookup words on different precompiled word lists.
This is the recommended mode of operation by the inventor of **diceware**
himself. An example word list would look like this:

    $> diceware 12345 22334 44556 54321 34251
      apathy  dang  oxeye  slain  issue

In the absence of physical dice you can also use **diceware** to generate a
passphrase by rolling electronic dice. An example word list made by the
-electronic flag would look like this:

    $> diceware -electronic
      hondo  groton  tabula  pete  2001  cock


## Usage:

    Usage:
    diceware [opts] [roll1] [roll2] [roll3] [...]

    Options:
      -dump
            dump the content of a -list, combine with -horizontal, -verbose
      -electronic
            roll dice electronically, see diceware FAQ
      -extra
            modify one word according to 'extra' rules (-electronic)
      -file string
            read word list from file
      -horizontal
            list rolled dice horizontally (-electronic) (default true)
      -list string
            name of list to use, see -lists (default "diceware")
      -lists
            list internal lists
      -rolls int
            number of rolls for -electronic (default 6)
      -verbose
            be more verbose (print line number of used word)
      -version
            show version, combine with -version



## Building:

    $> export GOPATH=`pwd`
    $> go get -v github.com/mgumz/diceware
    $> cp bin/diceware /usr/local/bin

## Integrated word lists:

* diceware - the original word list made by Arnold Reinhold
* diceware8k - the original 8k word list, made by Arnhold Reinhold
* beale - the alternative word list made by Alan Beale
* eff2016 - an alternative list by the Electronic Frontier Foundation

* ca - Catalan list, compiled by Marcel Hernandez (CC-BY-4.0)
* da - Danish list, compiled by Povl Falk-Jensen (CC-BY-4.0)
* de - German list, compiled by Benjamin Tenne (GPL2 license)
* eo - Esperanto list, compiled by Makis Diras
* es - Spanish list, compiled by Manuel Palao (GFDL)
* fi - Finish list, compiled by Martin Vermeer, Pauli Virtanen (GPL)
* fr - French list, compiled by Matthieu Weber (GPL)
* it - Italian list, compiled by Tarin Gamberini (GPL)
* jp - Japanese list, compiled by Hiroshi Yuki
* mi - Maori list, compiled by Rangi Kemara (CC-BY)
* nl - Dutch list, compiled by Bart Van den Eynde (GPL)
* no - Norwegian list, compiled by Willy T. Koch (CC-BY)
* pl - Polish list, compiled by Piotr (DrFugazi) Tarnowski
* pt - Portugese list, compiled by Patxi Pierce
* ru - compiled by "kitten"
* sv - Svedish list, compiled by Magnus Bodin
* sv8k - Svedish 8k list, compiled by Magnus Bodin
* tr - Turkish list, compiled by Mert Dirik

The author of **diceware** did not create any of the integrated lists, all
the effort, fame, blame etc belongs to the fine human beings who created
them. The lists are integrated only to increase the convenience of having
"all" lists at one place.


## Similar tools:

* https://github.com/natefinch/diceware/


## License:

Copyright (c) 2016 Mathias Gumz. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
     notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
     copyright notice, this list of conditions and the following
     disclaimer in the documentation and/or other materials provided
     with the distribution.
   * Neither the name of Mathias Gumz nor the names of diceware's
     contributors may be used to endorse or promote products derived
     from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

This License does not covers or alters the license of each included word
list.


[diceware]: http://world.std.com/~reinhold/diceware.html
[xkcd]: https://xkcd.com/936/
