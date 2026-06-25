ifeq ($(OS), Windows_NT)
    LIBPRZETAK = przetak.dll
    LIBPRZETAK_STATIC = przetak.lib
else
    UNAME := $(shell uname)
    ifeq ($(UNAME), Darwin)
        LIBPRZETAK = libprzetak.dylib
    else
        LIBPRZETAK = libprzetak.so
    endif
    LIBPRZETAK_STATIC = libprzetak.a
endif

$(LIBPRZETAK): przetak.go go/przetak.go go/translit.go go/abusive.go go/vulgar_n.go go/vulgar_p.go
	go build -buildmode=c-shared -o $(LIBPRZETAK) przetak.go

# Used for the rust build
$(LIBPRZETAK_STATIC): przetak.go go/przetak.go go/translit.go go/abusive.go go/vulgar_n.go go/vulgar_p.go
	go build -buildmode=c-archive -o $(LIBPRZETAK_STATIC) przetak.go

.PHONY: static
static: $(LIBPRZETAK_STATIC)
