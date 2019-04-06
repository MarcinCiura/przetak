/*
Copyright 2019 Marcin Ciura

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package main implements a dynamically linked library for checking
// whether UTF-8 strings contain abusive or vulgar speech in Polish.
package main

import (
	"C"
	"github.com/mciura/przetak/go"
	"reflect"
	"unsafe"
)

//export version
// version returns the date of last modification of package przetak
// in YYYY-MM-DD format.
func version() string {
	return przetak.VERSION
}

//export evaluate
// evaluate returns an integer whose bits with values 1, 2, or 4
// are set if the UTF-8 string s contains abusive words, vulgar
// words with negative connotations, or vulgar words with positive
// connotations in Polish, respectively.
func evaluate(s string) int32 {
	return przetak.Evaluate(s)
}

//export evaluatePP
// evaluatePP saves to res an integer whose bits with values 1, 2,
// or 4 are set if the UTF-8 string of length n pointed to by p
// contains abusive words, vulgar words with negative connotations,
// or vulgar words with positive connotations in Polish, respectively.
//
// It exists only because of Perl and R. The Perl FFI does not allow
// passing structs (Go strings) by value. The R FFI does not allow
// any method of passing structs, and its calling convention passes
// everything by reference.
func evaluatePP(p *unsafe.Pointer, n *int32, res *int32) {
	sh := reflect.StringHeader{Data: uintptr(*p), Len: int(*n)}
	*res = przetak.Evaluate(*(*string)(unsafe.Pointer(&sh)))
}

func main() {}
