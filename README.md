# *diceware*

**diceware** creates passphrases along the guidelines posted on the [diceware]
page and [xkcd]. In short, **diceware** makes use of physical dice to select
words out of word lists. a word list consists of 7776 words, each word yields
~12.9 bits of entropy. Thus, a passphrase made out of 8 words gives a whopping
103.2 bits of entropy + the advantage of having something easy to remember.

**diceware** can be used to lookup words on different precompiled word lists.
This is the recommended mode of operation by the inventor of **diceware**
himself.

In the absence of physical dice you can also use **diceware** to generate a
passphrase by rolling electronic dice.

## Usage:

    Usage:
    diceware [opts] [roll1] [roll2] [roll3] [...]

    Options:
      -dump
          dump the content of a -list
      -electronic
          roll dice electronically (see diceware FAQ)
      -extra
          modify one word according to 'extra' rules (-electronic)
      -file string
          read word list from file
      -horizontal
          list rolled dice horizontally (-electronic) (default true)
      -list string
          name of list to use (default "diceware")
      -lists
          list internal lists
      -rolls int
          number of rolls for -electronic (default 6)
      -verbose
          be more verbose (print line number of used word)


## Building:

    $> export GOPATH=`pwd`
    $> go get -v github.com/mgumz/diceware
    $> cp bin/diceware /usr/local/bin

## Compiled in word lists:

* diceware - the original word list
* beale - the alternative word list made by Alan Beale
* german - compiled by Benjamin Tenne (GPL2 license)
* russian - compiled by "kitten"

## Todo:

* optionally support *diceware8k*
* more built in word lists

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
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

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


[diceware]: http://world.std.com/~reinhold/diceware.html
[xkcd]: https://xkcd.com/936/
