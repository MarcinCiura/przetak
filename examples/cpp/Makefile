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

example: example.cpp $(LIBPRZETAK_H)
	g++ -Os -o example example.cpp $(LIBPRZETAK)
