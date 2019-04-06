# Examples

In the subdirectories, you can find example programs using Przetak
and instructions how to run them. All the examples, except the Go
program, use FFI (Foreign Function Interface) to call functions from
a dynamically linked library.

Each program reads one line and outputs its evaluation:

* 0 if the line contains neither abusive nor vulgar speech in Polish,
* 1 if it contains abusive speech,
* 2 if it contains vulgar speech with negative connotations,
* 4 if it contains vulgar speech with positive connotations,
* Binary OR of 1, 2, or 4 if it contains several kinds of toxic speech.
