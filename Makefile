ifeq ($(OS), Windows_NT)
    LIBPRZETAK = przetak.dll
else
    UNAME := $(shell uname)
    ifeq ($(UNAME), Darwin)
        LIBPRZETAK = libprzetak.dylib
    else
        LIBPRZETAK = libprzetak.so
    endif
endif

$(LIBPRZETAK): przetak.go go/przetak.go go/translit.go go/abusive.go go/vulgar_n.go go/vulgar_p.go
	go build -buildmode=c-shared -o $(LIBPRZETAK) przetak.go
