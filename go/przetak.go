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

// Package przetak implements a library for checking whether
// UTF-8 strings contain abusive or vulgar speech in Polish.
package przetak

import (
	"golang.org/x/text/unicode/norm"
)

const (
	// Three kinds of toxic speech.
	ABUSIVE  = int32(1)
	VULGAR_N = int32(2)
	VULGAR_P = int32(4)
	// Date of last modification of the package.
	VERSION = "2018-04-01"
)

// normalize returns Unicode Normalization Form KC (Compatibility
// Decomposition followed by Canonical Composition) of s.
//
// For instance, it replaces both \u212b (Angstrom sign) and A\u030a
// (A + combining ring above) with \u00e5 (Latin capital letter A
// with ring above).
func normalize(s string) string {
	return norm.NFKC.String(s)
}

// letterize returns a slice of runes corresponding to s, replacing
// Basic Latin digits, Latin-1 Supplement letters, Latin Extended-A
// letters, Latin Extended-B letters, IPA Extensions, Latin Extended
// Additional letters, Greek and Coptic letters, Cyrillic letters,
// and Cyrillic supplement letters with similar Latin lowercase
// letters, punctuation and symbols with periods, and the remaining
// characters with spaces.
//
// For instance, it converts both \u00e5 (Latin capital letter A with
// ring above), \u0391 (Greek capital letter Alpha), and \u0410
// (Cyrillic capital letter A) to 97 (Latin small letter a).
func letterize(s string) []rune {
	var ret []rune
	for _, r := range s {
		if r < kCyrillicEnd1 {
			ret = append(ret, kTranslitA[r])
		} else if kVietnameseStart <= r && r < kVietnameseEnd1 {
			ret = append(ret, kTranslitB[r-kVietnameseStart])
		} else {
			ret = append(ret, ' ')
		}
	}
	return ret
}

// spaceInHelper compresses s[:en] to s[:j] by removing spaces
// and cleans up the leftovers by filling s[j:] with spaces.
func spaceInHelper(s []rune, en int) {
	j := 0
	for _, r := range s[:en] {
		if r > ' ' {
			s[j] = r
			j++
		}
	}
	for i := range s[j:] {
		s[j+i] = ' '
	}
}

// spaceIn joins spaced out parts of s. Eligible for joining are
// single letters, immediately preceded or followed by optional
// sequences of periods and separated by sequences of one or more
// periods or spaces.
//
// For instance, it replaces []rune("f.o  .o.. bar") with
// []rune("f.o.o..    bar").
func spaceIn(s []rune) {
	var bg, en int
	state := 0
	for i, r := range s {
		switch state {
		// The finite-state machine corresponds to this regexp:
		// 0 1   2   1    2    3    4   5    4
		// `[a]([.]+[a])*[.]*([ ]+([.]*[a])*[.]*)+`
		case -1: // [a][a.]+
			if r == ' ' {
				state = 0
			}
		case 0: // [ .]
			if r > '.' {
				state = 1
			}
		case 1: // [a]
			switch {
			case r == ' ':
				bg = i
				en = i
				state = 3
			case r == '.':
				state = 2
			case r > '.':
				state = -1
			}
		case 2: // [.]
			switch {
			case r == ' ':
				bg = i
				en = i
				state = 3
			case r > '.':
				state = 1
			}
		case 3: // [ ]
			switch {
			case r == '.':
				state = 4
			case r > '.':
				state = 5
			}
		case 4: // [.]
			switch {
			case r == ' ':
				en = i
				state = 3
			case r > '.':
				state = 5
			}
		case 5: // [a]
			switch {
			case r == ' ':
				en = i
				state = 3
			case r == '.':
				state = 4
			case r > '.':
				spaceInHelper(s[bg:i-1], en-bg)
				state = -1
			}
		}
	}
	if 3 <= state {
		spaceInHelper(s[bg:], len(s)-bg)
	}
}

// forAllWords calls f(w, result), where w are sequences of letters
// from s. In addition, it calls f(jw, result), where jw are joined
// sequences of letters that were separated by sequences of periods
// in s. It returns early if f returns true.
//
// For instance, called with s = []rune("foo.bar"), it calls
// f([]rune("foo"), result), f([]rune("bar", result),
// and f([]rune("foobar", result) unless some call to f returns true.
func forAllWords(s []rune, f func([]rune, *int32) bool, result *int32) {
	var joined []rune
	var bg, en int
	state := 0
	for i, r := range s {
		switch state {
		// The finite-state machine corresponds to this regexp:
		// 0 1   2    3   4
		// `[a]+[.]+([a]+[.]+)*`
		case 0: // [ .]
			if r > '.' {
				bg = i
				state = 1
			}
		case 1: // [a]
			switch {
			case r == ' ':
				if f(s[bg:i], result) {
					return
				}
				state = 0
			case r == '.':
				if f(s[bg:i], result) {
					return
				}
				en = i
				state = 2
			}
		case 2: // [.]
			switch {
			case r == ' ':
				state = 0
			case r > '.':
				joined = append([]rune(nil), s[bg:en]...)
				bg = i
				state = 3
			}
		case 3: // [a]
			switch {
			case r == ' ':
				if f(s[bg:i], result) {
					return
				}
				joined = append(joined, s[bg:i]...)
				if f(joined, result) {
					return
				}
				state = 0
			case r == '.':
				if f(s[bg:i], result) {
					return
				}
				joined = append(joined, s[bg:i]...)
				state = 4
			}
		case 4: // [.]
			switch {
			case r == ' ':
				if f(joined, result) {
					return
				}
				state = 0
			case r > '.':
				bg = i
				state = 3
			}
		}
	}
	switch state {
	case 1:
		f(s[bg:], result)
	case 3:
		if f(s[bg:], result) {
			return
		}
		joined = append(joined, s[bg:]...)
		f(joined, result)
	case 4:
		f(joined, result)
	}
}

// dereplicate removes repeated runes from w.
func dereplicate(w []rune) []rune {
	j := 0
	for i, r := range w {
		if i == 0 || r != w[i-1] {
			w[j] = r
			j++
		}
	}
	return w[:j]
}

// evaluator contains three maps with 5-gram scores.
type evaluator struct {
	// The runes are obfuscated by XOR-ing them with 2.
	abusiveScore map[[5]rune]float32
	vulgarNScore map[[5]rune]float32
	vulgarPScore map[[5]rune]float32
}

// evaluateWord sets the ABUSIVE, VULGAR_N, or VULGAR_P bit in
// *result if w is an abusive word, vulgar word with negative
// connotations, or vulgar word with positive connotations in
// Polish, respectively. It returns true if *result has all
// three bits set.
func (e *evaluator) evaluateWord(w []rune, result *int32) bool {
	w = dereplicate(w)
	sl := [5]rune{0, 0, 0, 0, 0}
	abusive := e.abusiveScore[sl]
	vulgarN := e.vulgarNScore[sl]
	vulgarP := e.vulgarPScore[sl]
	switch l := len(w); l {
	case 0:
		panic("Empty word")
	case 1:
		sl := [5]rune{0, w[0] ^ 2, 0, 0, 0}
		abusive += e.abusiveScore[sl]
		vulgarN += e.vulgarNScore[sl]
		vulgarP += e.vulgarPScore[sl]
	case 2:
		sl := [5]rune{0, w[0] ^ 2, w[1] ^ 2, 0, 0}
		abusive += e.abusiveScore[sl]
		vulgarN += e.vulgarNScore[sl]
		vulgarP += e.vulgarPScore[sl]
	case 3:
		sl := [5]rune{0, w[0] ^ 2, w[1] ^ 2, w[2] ^ 2, 0}
		abusive += e.abusiveScore[sl]
		vulgarN += e.vulgarNScore[sl]
		vulgarP += e.vulgarPScore[sl]
	default:
		sl := [5]rune{0, w[0] ^ 2, w[1] ^ 2, w[2] ^ 2, w[3] ^ 2}
		abusive += e.abusiveScore[sl]
		vulgarN += e.vulgarNScore[sl]
		vulgarP += e.vulgarPScore[sl]
		for i, r := range w[5-1:] {
			sl := [5]rune{w[i] ^ 2, w[i+1] ^ 2, w[i+2] ^ 2, w[i+3] ^ 2, r ^ 2}
			abusive += e.abusiveScore[sl]
			vulgarN += e.vulgarNScore[sl]
			vulgarP += e.vulgarPScore[sl]
		}
		sl = [5]rune{w[l-4] ^ 2, w[l-3] ^ 2, w[l-2] ^ 2, w[l-1] ^ 2, 0}
		abusive += e.abusiveScore[sl]
		vulgarN += e.vulgarNScore[sl]
		vulgarP += e.vulgarPScore[sl]
	}
	if abusive > 0 {
		*result |= ABUSIVE
	}
	if vulgarN > 0 {
		*result |= VULGAR_N
	}
	if vulgarP > 0 {
		*result |= VULGAR_P
	}
	return (*result == ABUSIVE|VULGAR_N|VULGAR_P)
}

// evaluate returns a bit mask whose ABUSIVE, VULGAR_N, or
// VULGAR_P bits are set if s contains abusive words, vulgar
// words with negative connotations, or vulgar words with
// positive connotations in Polish, respectively.
func (e *evaluator) evaluate(s string) int32 {
	n := normalize(s)
	r := letterize(n)
	spaceIn(r)
	result := int32(0)
	forAllWords(r, e.evaluateWord, &result)
	return result
}

// kEvaluator is the global evaluator.
var kEvaluator = evaluator{
	abusiveScore: kAbusiveScore,
	vulgarNScore: kVulgarNScore,
	vulgarPScore: kVulgarPScore,
}

// Evaluate returns a bit mask whose ABUSIVE, VULGAR_N, and
// VULGAR_P bits are set if s contains abusive words, vulgar
// words with negative connotations, or vulgar words with
// positive connotations in Polish, respectively.
func Evaluate(s string) int32 {
	return kEvaluator.evaluate(s)
}
