# Java

To install dependencies, compile the program, and run it on Ubuntu:

```
$ sudo apt install libjna-java
$ make
$ java -cp .:../..:/usr/share/java/jna.jar Example
```

Other Linux distributions or operating systems may require changes
to the first line that installs dynamic access of native libraries
from Java without JNI (Java Native Interface), and to the
`/usr/share/java/jna.jar` path in `Makefile` and in the third line.
