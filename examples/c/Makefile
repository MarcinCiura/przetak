ifeq ($(OS), Windows_NT)
    LIBPRZETAK = ../../przetak.dll
    LIBPRZETAK_H = ../../przetak.h
else
    UNAME := $(shell uname)
    ifeq ($(UNAME), Darwin)
        LIBPRZETAK = ../../libprzetak.dylib
    else
        LIBPRZETAK = ../../libprzetak.so
    endif
    LIBPRZETAK_H = ../../libprzetak.h
endif

example: example.c $(LIBPRZETAK_H)
	gcc -Os -o example example.c $(LIBPRZETAK)
