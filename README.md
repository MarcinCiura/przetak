# Przetak: fewer weeds on the Web

Przetak is a library for checking whether a text contains
abusive or vulgar speech in Polish. While it is written in Go,
it can be used by programs written in many other languages
thanks to FFI (Foreign Function Interface).

Przetak is resilient to:

* replicating letters,
* spacing out the words,
* inserting non-letters between letters,
* homograph spoofing, i.e. replacing letters with similar characters.

Also, thanks to its use of character 5-grams, it handles some
frequent misspellings and out-of-vocabulary words composed of
morphemes with an abusive or vulgar meaning.

Przetak finished the Polish contest of cyberbullying detection
PolEval 2019 [in second place](http://poleval.pl/index.php/results/).
[Here](http://poleval.pl/files/poleval2019.pdf#page=127) is
a paper about Przetak, and [here](http://2019.poleval.pl/files/2019/15.pdf)
are the slides from my presentation at AI & NLP Workshop Day 2019.

## Installation

First, get the package:

```
$ go get github.com/mciura/przetak
```

Change directory to your `${GOPATH}/src/github.com/mciura/przetak`
and run `make` to build the shared library. Depending on your
operating system, the shared library will be called:

* `libprzetak.so` on Linux,
* `libprzetak.dylib` on macOS,
* `przetak.dll` on Windows.

## Usage

Przetak's `evaluate()` function returns an integer whose
bits with respective values 1, 2, or 4 are set if the input
UTF-8 string contains:

* abusive words, e.g. _oszołom_,
* vulgar words with negative connotations, e.g. _ku**a_,
* vulgar words with positive connotations, e.g. _za**biście_.

The [examples](examples)
directory showcases the use of Przetak directly from Go
and from several other programming languages via FFI
(Foreign Function Interface).

## Author

Marcin Ciura < mciura at gmail dot com >

## License

Przetak is licensed under
[Apache License, Version 2.0](LICENSE).
