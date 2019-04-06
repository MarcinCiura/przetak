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

package przetak

import (
	"bytes"
	"testing"
)

func TestNormalize(t *testing.T) {
	const kelvin = "\u212a"
	const combOgonek = "\u0328"
	const combDotAbove = "\u0307"
	const input = kelvin + "sia" + combOgonek + "z" + combDotAbove + "ka"

	actual := normalize(input)
	const expected = "Książka"
	if actual != expected {
		t.Errorf("Expected `%v`; got `%v`", expected, actual)
	}
}

func runeSlicesEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestLetterize(t *testing.T) {
	const hiraganaA = "\u3042"
	const cyrillicK = "\u043a"
	const vietnameseA = "\u1e01"
	const input = hiraganaA + ", K5iąż" + cyrillicK + vietnameseA

	actual := letterize(input)
	expected := []rune(" . książka")
	if !runeSlicesEqual(actual, expected) {
		t.Errorf("Expected `%v`; got `%v`",
			string(expected), string(actual))

	}
}

func TestSpaceIn(t *testing.T) {
	text := []rune(" . xx ..s. .p a . .c. .e.. d out S.P A C E D")
	spaceIn(text)
	expected := []rune(" . xx ..s..pa..c..e..d       out S.PACED    ")
	if !runeSlicesEqual(text, expected) {
		t.Errorf("Expected `%v`; got `%v`",
			string(expected), string(text))
	}
}

func TestSpaceInEndAtPunctuation(t *testing.T) {
	text := []rune("a  b  c.")
	spaceIn(text)
	expected := []rune("abc.    ")
	if !runeSlicesEqual(text, expected) {
		t.Errorf("Expected `%v`; got `%v`",
			string(expected), string(text))
	}
}

func TestSpaceInEndAtSpace(t *testing.T) {
	text := []rune("a b c ")
	spaceIn(text)
	expected := []rune("abc   ")
	if !runeSlicesEqual(text, expected) {
		t.Errorf("Expected `%v`; got `%v`",
			string(expected), string(text))
	}
}

func TestSpaceInSingleCharInBetween(t *testing.T) {
	text := []rune("ten i ów")
	spaceIn(text)
	expected := []rune("ten i ów")
	if !runeSlicesEqual(text, expected) {
		t.Errorf("Expected `%v`; got `%v`",
			string(expected), string(text))
	}
}

func TestSpaceInTwoSingleChars(t *testing.T) {
	text := []rune("a i ów")
	spaceIn(text)
	expected := []rune("ai  ów")
	if !runeSlicesEqual(text, expected) {
		t.Errorf("Expected `%v`; got `%v`",
			string(expected), string(text))
	}
}

func runeSliceSlicesEqual(a, b [][]rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, s := range a {
		if !runeSlicesEqual(s, b[i]) {
			return false
		}
	}
	return true
}

func runeSliceSliceToString(rss [][]rune) string {
	ret := bytes.NewBuffer([]byte(`"`))
	for i, rs := range rss {
		if i > 0 {
			ret.WriteString(`", "`)
		}
		for _, r := range rs {
			ret.WriteRune(r)
		}
	}
	ret.WriteString(`"`)
	return ret.String()
}

func TestForAllWords(t *testing.T) {
	input := []rune("x w1. u2.v3 u4.v5. u6.v7.w8.")
	var actual [][]rune
	var dummy int32
	forAllWords(input, func(word []rune, result *int32) bool {
		actual = append(actual, word)
		return false
	}, &dummy)
	expected := [][]rune{
		[]rune("x"),
		[]rune("w1"),
		[]rune("u2"),
		[]rune("v3"),
		[]rune("u2v3"),
		[]rune("u4"),
		[]rune("v5"),
		[]rune("u4v5"),
		[]rune("u6"),
		[]rune("v7"),
		[]rune("w8"),
		[]rune("u6v7w8"),
	}
	if !runeSliceSlicesEqual(actual, expected) {
		t.Errorf("Expected [%v]; got [%v]",
			runeSliceSliceToString(expected),
			runeSliceSliceToString(actual))
	}
}

func TestForAllWordsEndAtWord1(t *testing.T) {
	input := []rune("u1")
	var actual [][]rune
	var dummy int32
	forAllWords(input, func(word []rune, result *int32) bool {
		actual = append(actual, word)
		return false
	}, &dummy)
	expected := [][]rune{
		[]rune("u1"),
	}
	if !runeSliceSlicesEqual(actual, expected) {
		t.Errorf("Expected [%v]; got [%v]",
			runeSliceSliceToString(expected),
			runeSliceSliceToString(actual))
	}
}

func TestForAllWordsEndAtWord2(t *testing.T) {
	input := []rune("u1.v2")
	var actual [][]rune
	var dummy int32
	forAllWords(input, func(word []rune, result *int32) bool {
		actual = append(actual, word)
		return false
	}, &dummy)
	expected := [][]rune{
		[]rune("u1"),
		[]rune("v2"),
		[]rune("u1v2"),
	}
	if !runeSliceSlicesEqual(actual, expected) {
		t.Errorf("Expected [%v]; got [%v]",
			runeSliceSliceToString(expected),
			runeSliceSliceToString(actual))
	}
}

func TestForAllWordsEarlyExit(t *testing.T) {
	input := []rune("x w1. u2.v3 u4.v5. u6.v7.w8.")
	var actual [][]rune
	var result int32
	forAllWords(input, func(word []rune, result *int32) bool {
		actual = append(actual, word)
		*result = int32(len(actual))
		return (*result >= 3)
	}, &result)
	expected := [][]rune{
		[]rune("x"),
		[]rune("w1"),
		[]rune("u2"),
	}
	if !runeSliceSlicesEqual(actual, expected) {
		t.Errorf("Expected [%v]; got [%v]",
			runeSliceSliceToString(expected),
			runeSliceSliceToString(actual))
	}
	if result != 3 {
		t.Errorf("Expected 3; got %v", result)
	}
}

func TestDereplicate(t *testing.T) {
	input := []rune("inną")
	actual := dereplicate(input)
	expected := []rune("iną")
	if !runeSlicesEqual(actual, expected) {
		t.Errorf("Expected [%v]; got [%v]",
			string(expected), string(actual))
	}
}

func TestEvaluate(t *testing.T) {
	var test_score = map[[5]rune]float32{
		[5]rune{'k' ^ 2, 'w' ^ 2, 'i' ^ 2, 'a' ^ 2, 't' ^ 2}: 2.0,
		[5]rune{0, 0, 0, 0, 0}:                               -1.0,
	}
	var test_evaluator = evaluator{
		abusiveScore: test_score,
		vulgarNScore: test_score,
		vulgarPScore: map[[5]rune]float32{},
	}
	actual := test_evaluator.evaluate("nic")
	if actual != 0 {
		t.Errorf("Expected 0; got %v", actual)
	}

	const kelvin = "\u212a"
	actual = test_evaluator.evaluate(kelvin + "*w**i***aa****t")
	expected := ABUSIVE | VULGAR_N
	if actual != expected {
		t.Errorf("Expected %v; got %v", expected, actual)
	}

	actual = test_evaluator.evaluate("k  w  i  a  t  e  k")
	if actual != expected {
		t.Errorf("Expected %v; got %v", expected, actual)
	}
}
