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

const (
	kCyrillicEnd1    = 0x0530
	kVietnameseStart = 0x1e00
	kVietnameseEnd1  = 0x1f00
)

var kTranslitA = [kCyrillicEnd1]rune{
	0x000: ' ',
	0x001: ' ',
	0x002: ' ',
	0x003: ' ',
	0x004: ' ',
	0x005: ' ',
	0x006: ' ',
	0x007: ' ',
	0x008: ' ',
	0x009: ' ',
	0x00a: ' ',
	0x00b: ' ',
	0x00c: ' ',
	0x00d: ' ',
	0x00e: ' ',
	0x00f: ' ',
	0x010: ' ',
	0x011: ' ',
	0x012: ' ',
	0x013: ' ',
	0x014: ' ',
	0x015: ' ',
	0x016: ' ',
	0x017: ' ',
	0x018: ' ',
	0x019: ' ',
	0x01a: ' ',
	0x01b: ' ',
	0x01c: ' ',
	0x01d: ' ',
	0x01e: ' ',
	0x01f: ' ',
	' ':   ' ',
	'!':   '.',
	'"':   '.',
	'#':   '.',
	'$':   '.',
	'%':   '.',
	'&':   '.',
	'\'':  '.',
	'(':   '.',
	')':   '.',
	'*':   '.',
	'+':   '.',
	',':   '.',
	'-':   '.',
	'.':   '.',
	'/':   '.',
	'0':   'o',
	'1':   'i',
	'2':   'z',
	'3':   'e',
	'4':   'a',
	'5':   's',
	'6':   'b',
	'7':   't',
	'8':   'b',
	'9':   'g',
	':':   '.',
	';':   '.',
	'<':   '.',
	'=':   '.',
	'>':   '.',
	'?':   '.',
	'@':   '.',
	'A':   'a',
	'B':   'b',
	'C':   'c',
	'D':   'd',
	'E':   'e',
	'F':   'f',
	'G':   'g',
	'H':   'h',
	'I':   'i',
	'J':   'j',
	'K':   'k',
	'L':   'l',
	'M':   'm',
	'N':   'n',
	'O':   'o',
	'P':   'p',
	'Q':   'q',
	'R':   'r',
	'S':   's',
	'T':   't',
	'U':   'u',
	'V':   'v',
	'W':   'w',
	'X':   'x',
	'Y':   'y',
	'Z':   'z',
	'[':   '.',
	'\\':  '.',
	']':   '.',
	'^':   '.',
	'_':   '.',
	'`':   '.',
	'a':   'a',
	'b':   'b',
	'c':   'c',
	'd':   'd',
	'e':   'e',
	'f':   'f',
	'g':   'g',
	'h':   'h',
	'i':   'i',
	'j':   'j',
	'k':   'k',
	'l':   'l',
	'm':   'm',
	'n':   'n',
	'o':   'o',
	'p':   'p',
	'q':   'q',
	'r':   'r',
	's':   's',
	't':   't',
	'u':   'u',
	'v':   'v',
	'w':   'w',
	'x':   'x',
	'y':   'y',
	'z':   'z',
	'{':   '.',
	'|':   '.',
	'}':   '.',
	'~':   '.',
	0x07f: ' ',
	0x080: ' ',
	0x081: ' ',
	0x082: ' ',
	0x083: ' ',
	0x084: ' ',
	0x085: ' ',
	0x086: ' ',
	0x087: ' ',
	0x088: ' ',
	0x089: ' ',
	0x08a: ' ',
	0x08b: ' ',
	0x08c: ' ',
	0x08d: ' ',
	0x08e: ' ',
	0x08f: ' ',
	0x090: ' ',
	0x091: ' ',
	0x092: ' ',
	0x093: ' ',
	0x094: ' ',
	0x095: ' ',
	0x096: ' ',
	0x097: ' ',
	0x098: ' ',
	0x099: ' ',
	0x09a: ' ',
	0x09b: ' ',
	0x09c: ' ',
	0x09d: ' ',
	0x09e: ' ',
	0x09f: ' ',
	0x0a0: ' ',
	'¡':   'i',
	'¢':   'c',
	'£':   'l',
	'¤':   'o',
	'¥':   'y',
	'¦':   'l',
	'§':   's',
	'¨':   '.',
	'©':   'c',
	'ª':   'a',
	'«':   '.',
	'¬':   '.',
	0x0ad: ' ',
	'®':   'r',
	'¯':   '.',
	'°':   '.',
	'±':   '.',
	'²':   '.',
	'³':   '.',
	'´':   '.',
	'µ':   '.',
	'¶':   '.',
	'·':   '.',
	'¸':   '.',
	'¹':   '.',
	'º':   'o',
	'»':   '.',
	'¼':   '.',
	'½':   '.',
	'¾':   '.',
	'¿':   '.',
	'À':   'a',
	'Á':   'a',
	'Â':   'a',
	'Ã':   'a',
	'Ä':   'a',
	'Å':   'a',
	'Æ':   'e',
	'Ç':   'c',
	'È':   'e',
	'É':   'e',
	'Ê':   'e',
	'Ë':   'e',
	'Ì':   'i',
	'Í':   'i',
	'Î':   'i',
	'Ï':   'i',
	'Ð':   'd',
	'Ñ':   'n',
	'Ò':   'o',
	'Ó':   'ó',
	'Ô':   'o',
	'Õ':   'o',
	'Ö':   'o',
	'×':   '.',
	'Ø':   'o',
	'Ù':   'u',
	'Ú':   'u',
	'Û':   'u',
	'Ü':   'u',
	'Ý':   'y',
	'Þ':   'p',
	'ß':   'b',
	'à':   'a',
	'á':   'a',
	'â':   'a',
	'ã':   'a',
	'ä':   'a',
	'å':   'a',
	'æ':   'e',
	'ç':   'c',
	'è':   'e',
	'é':   'e',
	'ê':   'e',
	'ë':   'e',
	'ì':   'i',
	'í':   'i',
	'î':   'i',
	'ï':   'i',
	'ð':   'd',
	'ñ':   'n',
	'ò':   'o',
	'ó':   'ó',
	'ô':   'o',
	'õ':   'o',
	'ö':   'o',
	'÷':   '.',
	'ø':   'o',
	'ù':   'u',
	'ú':   'u',
	'û':   'u',
	'ü':   'u',
	'ý':   'y',
	'þ':   'p',
	'ÿ':   'y',
	'Ā':   'a',
	'ā':   'a',
	'Ă':   'a',
	'ă':   'a',
	'Ą':   'ą',
	'ą':   'ą',
	'Ć':   'ć',
	'ć':   'ć',
	'Ĉ':   'c',
	'ĉ':   'c',
	'Ċ':   'c',
	'ċ':   'c',
	'Č':   'c',
	'č':   'c',
	'Ď':   'd',
	'ď':   'd',
	'Đ':   'd',
	'đ':   'd',
	'Ē':   'e',
	'ē':   'e',
	'Ĕ':   'e',
	'ĕ':   'e',
	'Ė':   'e',
	'ė':   'e',
	'Ę':   'ę',
	'ę':   'ę',
	'Ě':   'e',
	'ě':   'e',
	'Ĝ':   'g',
	'ĝ':   'g',
	'Ğ':   'g',
	'ğ':   'g',
	'Ġ':   'g',
	'ġ':   'g',
	'Ģ':   'g',
	'ģ':   'g',
	'Ĥ':   'h',
	'ĥ':   'h',
	'Ħ':   'h',
	'ħ':   'h',
	'Ĩ':   'i',
	'ĩ':   'i',
	'Ī':   'i',
	'ī':   'i',
	'Ĭ':   'i',
	'ĭ':   'i',
	'Į':   'i',
	'į':   'i',
	'İ':   'i',
	'ı':   'i',
	'Ĳ':   'y',
	'ĳ':   'y',
	'Ĵ':   'j',
	'ĵ':   'j',
	'Ķ':   'k',
	'ķ':   'k',
	'ĸ':   'k',
	'Ĺ':   'l',
	'ĺ':   'l',
	'Ļ':   'l',
	'ļ':   'l',
	'Ľ':   'l',
	'ľ':   'l',
	'Ŀ':   'l',
	'ŀ':   'l',
	'Ł':   'ł',
	'ł':   'ł',
	'Ń':   'ń',
	'ń':   'ń',
	'Ņ':   'n',
	'ņ':   'n',
	'Ň':   'n',
	'ň':   'n',
	'ŉ':   'n',
	'Ŋ':   'n',
	'ŋ':   'n',
	'Ō':   'o',
	'ō':   'o',
	'Ŏ':   'o',
	'ŏ':   'o',
	'Ő':   'o',
	'ő':   'o',
	'Œ':   'e',
	'œ':   'e',
	'Ŕ':   'r',
	'ŕ':   'r',
	'Ŗ':   'r',
	'ŗ':   'r',
	'Ř':   'r',
	'ř':   'r',
	'Ś':   'ś',
	'ś':   'ś',
	'Ŝ':   's',
	'ŝ':   's',
	'Ş':   's',
	'ş':   's',
	'Š':   's',
	'š':   's',
	'Ţ':   't',
	'ţ':   't',
	'Ť':   't',
	'ť':   't',
	'Ŧ':   't',
	'ŧ':   't',
	'Ũ':   'u',
	'ũ':   'u',
	'Ū':   'u',
	'ū':   'u',
	'Ŭ':   'u',
	'ŭ':   'u',
	'Ů':   'u',
	'ů':   'u',
	'Ű':   'u',
	'ű':   'u',
	'Ų':   'u',
	'ų':   'u',
	'Ŵ':   'w',
	'ŵ':   'w',
	'Ŷ':   'y',
	'ŷ':   'y',
	'Ÿ':   'y',
	'Ź':   'ź',
	'ź':   'ź',
	'Ż':   'ż',
	'ż':   'ż',
	'Ž':   'z',
	'ž':   'z',
	'ſ':   's',
	'ƀ':   'b',
	'Ɓ':   'b',
	'Ƃ':   'b',
	'ƃ':   'b',
	'Ƅ':   'b',
	'ƅ':   'b',
	'Ɔ':   'c',
	'Ƈ':   'c',
	'ƈ':   'c',
	'Ɖ':   'd',
	'Ɗ':   'd',
	'Ƌ':   'd',
	'ƌ':   'd',
	'ƍ':   'o',
	'Ǝ':   'e',
	'Ə':   'e',
	'Ɛ':   'e',
	'Ƒ':   'f',
	'ƒ':   'f',
	'Ɠ':   'g',
	'Ɣ':   'y',
	'ƕ':   'h',
	'Ɩ':   'i',
	'Ɨ':   'i',
	'Ƙ':   'k',
	'ƙ':   'k',
	'ƚ':   'l',
	'ƛ':   'a',
	'Ɯ':   'w',
	'Ɲ':   'n',
	'ƞ':   'n',
	'Ɵ':   'o',
	'Ơ':   'o',
	'ơ':   'o',
	'Ƣ':   'o',
	'ƣ':   'o',
	'Ƥ':   'p',
	'ƥ':   'p',
	'Ʀ':   'r',
	'Ƨ':   's',
	'ƨ':   's',
	'Ʃ':   'e',
	'ƪ':   'l',
	'ƫ':   't',
	'Ƭ':   't',
	'ƭ':   't',
	'Ʈ':   't',
	'Ư':   'u',
	'ư':   'u',
	'Ʊ':   'u',
	'Ʋ':   'v',
	'Ƴ':   'y',
	'ƴ':   'y',
	'Ƶ':   'z',
	'ƶ':   'z',
	'Ʒ':   'z',
	'Ƹ':   'e',
	'ƹ':   'e',
	'ƺ':   'z',
	'ƻ':   'z',
	'Ƽ':   's',
	'ƽ':   's',
	'ƾ':   's',
	'ƿ':   'p',
	'ǀ':   'i',
	'ǁ':   'i',
	'ǂ':   'i',
	'ǃ':   'i',
	'Ǆ':   'd',
	'ǅ':   'd',
	'ǆ':   'd',
	'Ǉ':   'l',
	'ǈ':   'l',
	'ǉ':   'l',
	'Ǌ':   'n',
	'ǋ':   'n',
	'ǌ':   'n',
	'Ǎ':   'a',
	'ǎ':   'a',
	'Ǐ':   'i',
	'ǐ':   'i',
	'Ǒ':   'o',
	'ǒ':   'o',
	'Ǔ':   'u',
	'ǔ':   'u',
	'Ǖ':   'u',
	'ǖ':   'u',
	'Ǘ':   'u',
	'ǘ':   'u',
	'Ǚ':   'u',
	'ǚ':   'u',
	'Ǜ':   'u',
	'ǜ':   'u',
	'ǝ':   'e',
	'Ǟ':   'a',
	'ǟ':   'a',
	'Ǡ':   'a',
	'ǡ':   'a',
	'Ǣ':   'e',
	'ǣ':   'e',
	'Ǥ':   'g',
	'ǥ':   'g',
	'Ǧ':   'g',
	'ǧ':   'g',
	'Ǩ':   'k',
	'ǩ':   'k',
	'Ǫ':   'o',
	'ǫ':   'o',
	'Ǭ':   'o',
	'ǭ':   'o',
	'Ǯ':   'z',
	'ǯ':   'z',
	'ǰ':   'j',
	'Ǳ':   'd',
	'ǲ':   'd',
	'ǳ':   'd',
	'Ǵ':   'g',
	'ǵ':   'g',
	'Ƕ':   'h',
	'Ƿ':   'p',
	'Ǹ':   'n',
	'ǹ':   'n',
	'Ǻ':   'a',
	'ǻ':   'a',
	'Ǽ':   'e',
	'ǽ':   'e',
	'Ǿ':   'o',
	'ǿ':   'o',
	'Ȁ':   'a',
	'ȁ':   'a',
	'Ȃ':   'a',
	'ȃ':   'a',
	'Ȅ':   'e',
	'ȅ':   'e',
	'Ȇ':   'e',
	'ȇ':   'e',
	'Ȉ':   'i',
	'ȉ':   'i',
	'Ȋ':   'i',
	'ȋ':   'i',
	'Ȍ':   'o',
	'ȍ':   'o',
	'Ȏ':   'o',
	'ȏ':   'o',
	'Ȑ':   'r',
	'ȑ':   'r',
	'Ȓ':   'r',
	'ȓ':   'r',
	'Ȕ':   'u',
	'ȕ':   'u',
	'Ȗ':   'u',
	'ȗ':   'u',
	'Ș':   's',
	'ș':   's',
	'Ț':   't',
	'ț':   't',
	'Ȝ':   'z',
	'ȝ':   'z',
	'Ȟ':   'h',
	'ȟ':   'h',
	'Ƞ':   'n',
	'ȡ':   'd',
	'Ȣ':   'o',
	'ȣ':   'o',
	'Ȥ':   'z',
	'ȥ':   'z',
	'Ȧ':   'a',
	'ȧ':   'a',
	'Ȩ':   'e',
	'ȩ':   'e',
	'Ȫ':   'o',
	'ȫ':   'o',
	'Ȭ':   'o',
	'ȭ':   'o',
	'Ȯ':   'o',
	'ȯ':   'o',
	'Ȱ':   'o',
	'ȱ':   'o',
	'Ȳ':   'y',
	'ȳ':   'y',
	'ȴ':   'l',
	'ȵ':   'n',
	'ȶ':   't',
	'ȷ':   'j',
	'ȸ':   'b',
	'ȹ':   'p',
	'Ⱥ':   'a',
	'Ȼ':   'c',
	'ȼ':   'c',
	'Ƚ':   'l',
	'Ⱦ':   't',
	'ȿ':   's',
	'ɀ':   'z',
	'Ɂ':   'p',
	'ɂ':   'p',
	'Ƀ':   'b',
	'Ʉ':   'u',
	'Ʌ':   'a',
	'Ɇ':   'e',
	'ɇ':   'e',
	'Ɉ':   'j',
	'ɉ':   'j',
	'Ɋ':   'q',
	'ɋ':   'q',
	'Ɍ':   'r',
	'ɍ':   'r',
	'Ɏ':   'y',
	'ɏ':   'y',
	'ɐ':   'a',
	'ɑ':   'a',
	'ɒ':   'a',
	'ɓ':   'b',
	'ɔ':   'c',
	'ɕ':   'c',
	'ɖ':   'd',
	'ɗ':   'd',
	'ɘ':   'e',
	'ə':   'e',
	'ɚ':   'e',
	'ɛ':   'e',
	'ɜ':   'e',
	'ɝ':   'e',
	'ɞ':   'b',
	'ɟ':   'f',
	'ɠ':   'g',
	'ɡ':   'g',
	'ɢ':   'g',
	'ɣ':   'y',
	'ɤ':   'v',
	'ɥ':   'h',
	'ɦ':   'h',
	'ɧ':   'h',
	'ɨ':   'i',
	'ɩ':   'i',
	'ɪ':   'i',
	'ɫ':   'l',
	'ɬ':   'l',
	'ɭ':   'l',
	'ɮ':   'b',
	'ɯ':   'w',
	'ɰ':   'w',
	'ɱ':   'm',
	'ɲ':   'n',
	'ɳ':   'n',
	'ɴ':   'n',
	'ɵ':   'o',
	'ɶ':   'e',
	'ɷ':   'w',
	'ɸ':   'p',
	'ɹ':   'r',
	'ɺ':   'r',
	'ɻ':   'r',
	'ɼ':   'r',
	'ɽ':   'r',
	'ɾ':   'r',
	'ɿ':   'r',
	'ʀ':   'r',
	'ʁ':   'r',
	'ʂ':   's',
	'ʃ':   's',
	'ʄ':   's',
	'ʅ':   's',
	'ʆ':   's',
	'ʇ':   't',
	'ʈ':   't',
	'ʉ':   'u',
	'ʊ':   'u',
	'ʋ':   'v',
	'ʌ':   'a',
	'ʍ':   'm',
	'ʎ':   'y',
	'ʏ':   'y',
	'ʐ':   'z',
	'ʑ':   'z',
	'ʒ':   'z',
	'ʓ':   'z',
	'ʔ':   'p',
	'ʕ':   'c',
	'ʖ':   'j',
	'ʗ':   'c',
	'ʘ':   'o',
	'ʙ':   'b',
	'ʚ':   'o',
	'ʛ':   'g',
	'ʜ':   'h',
	'ʝ':   'j',
	'ʞ':   'k',
	'ʟ':   'l',
	'ʠ':   'q',
	'ʡ':   'p',
	'ʢ':   'c',
	'ʣ':   'd',
	'ʤ':   'd',
	'ʥ':   'd',
	'ʦ':   's',
	'ʧ':   's',
	'ʨ':   'c',
	'ʩ':   'n',
	'ʪ':   's',
	'ʫ':   'z',
	'ʬ':   'w',
	'ʭ':   'c',
	'ʮ':   'y',
	'ʯ':   'y',
	'ʰ':   '.',
	'ʱ':   '.',
	'ʲ':   '.',
	'ʳ':   '.',
	'ʴ':   '.',
	'ʵ':   '.',
	'ʶ':   '.',
	'ʷ':   '.',
	'ʸ':   '.',
	'ʹ':   '.',
	'ʺ':   '.',
	'ʻ':   '.',
	'ʼ':   '.',
	'ʽ':   '.',
	'ʾ':   '.',
	'ʿ':   '.',
	'ˀ':   '.',
	'ˁ':   '.',
	'˂':   '.',
	'˃':   '.',
	'˄':   '.',
	'˅':   '.',
	'ˆ':   '.',
	'ˇ':   '.',
	'ˈ':   '.',
	'ˉ':   '.',
	'ˊ':   '.',
	'ˋ':   '.',
	'ˌ':   '.',
	'ˍ':   '.',
	'ˎ':   '.',
	'ˏ':   '.',
	'ː':   '.',
	'ˑ':   '.',
	'˒':   '.',
	'˓':   '.',
	'˔':   '.',
	'˕':   '.',
	'˖':   '.',
	'˗':   '.',
	'˘':   '.',
	'˙':   '.',
	'˚':   '.',
	'˛':   '.',
	'˜':   '.',
	'˝':   '.',
	'˞':   '.',
	'˟':   '.',
	'ˠ':   '.',
	'ˡ':   '.',
	'ˢ':   '.',
	'ˣ':   '.',
	'ˤ':   '.',
	'˥':   '.',
	'˦':   '.',
	'˧':   '.',
	'˨':   '.',
	'˩':   '.',
	'˪':   '.',
	'˫':   '.',
	'ˬ':   '.',
	'˭':   '.',
	'ˮ':   '.',
	'˯':   '.',
	'˰':   '.',
	'˱':   '.',
	'˲':   '.',
	'˳':   '.',
	'˴':   '.',
	'˵':   '.',
	'˶':   '.',
	'˷':   '.',
	'˸':   '.',
	'˹':   '.',
	'˺':   '.',
	'˻':   '.',
	'˼':   '.',
	'˽':   '.',
	'˾':   '.',
	'˿':   '.',
	0x300: '.',
	0x301: '.',
	0x302: '.',
	0x303: '.',
	0x304: '.',
	0x305: '.',
	0x306: '.',
	0x307: '.',
	0x308: '.',
	0x309: '.',
	0x30a: '.',
	0x30b: '.',
	0x30c: '.',
	0x30d: '.',
	0x30e: '.',
	0x30f: '.',
	0x310: '.',
	0x311: '.',
	0x312: '.',
	0x313: '.',
	0x314: '.',
	0x315: '.',
	0x316: '.',
	0x317: '.',
	0x318: '.',
	0x319: '.',
	0x31a: '.',
	0x31b: '.',
	0x31c: '.',
	0x31d: '.',
	0x31e: '.',
	0x31f: '.',
	0x320: '.',
	0x321: '.',
	0x322: '.',
	0x323: '.',
	0x324: '.',
	0x325: '.',
	0x326: '.',
	0x327: '.',
	0x328: '.',
	0x329: '.',
	0x32a: '.',
	0x32b: '.',
	0x32c: '.',
	0x32d: '.',
	0x32e: '.',
	0x32f: '.',
	0x330: '.',
	0x331: '.',
	0x332: '.',
	0x333: '.',
	0x334: '.',
	0x335: '.',
	0x336: '.',
	0x337: '.',
	0x338: '.',
	0x339: '.',
	0x33a: '.',
	0x33b: '.',
	0x33c: '.',
	0x33d: '.',
	0x33e: '.',
	0x33f: '.',
	0x340: '.',
	0x341: '.',
	0x342: '.',
	0x343: '.',
	0x344: '.',
	0x345: '.',
	0x346: '.',
	0x347: '.',
	0x348: '.',
	0x349: '.',
	0x34a: '.',
	0x34b: '.',
	0x34c: '.',
	0x34d: '.',
	0x34e: '.',
	0x34f: '.',
	0x350: '.',
	0x351: '.',
	0x352: '.',
	0x353: '.',
	0x354: '.',
	0x355: '.',
	0x356: '.',
	0x357: '.',
	0x358: '.',
	0x359: '.',
	0x35a: '.',
	0x35b: '.',
	0x35c: '.',
	0x35d: '.',
	0x35e: '.',
	0x35f: '.',
	0x360: '.',
	0x361: '.',
	0x362: '.',
	0x363: '.',
	0x364: '.',
	0x365: '.',
	0x366: '.',
	0x367: '.',
	0x368: '.',
	0x369: '.',
	0x36a: '.',
	0x36b: '.',
	0x36c: '.',
	0x36d: '.',
	0x36e: '.',
	0x36f: '.',
	'Ͱ':   'i',
	'ͱ':   'i',
	'Ͳ':   't',
	'ͳ':   't',
	'ʹ':   '.',
	'͵':   '.',
	'Ͷ':   'n',
	'ͷ':   'n',
	0x378: '.',
	0x379: '.',
	'ͺ':   '.',
	'ͻ':   'c',
	'ͼ':   'c',
	'ͽ':   'c',
	';':   '.',
	'Ϳ':   'j',
	0x380: '.',
	0x381: '.',
	0x382: '.',
	0x383: '.',
	'΄':   '.',
	'΅':   '.',
	'Ά':   'a',
	'·':   '.',
	'Έ':   'e',
	'Ή':   'h',
	'Ί':   'i',
	0x38b: '.',
	'Ό':   'o',
	0x38d: '.',
	'Ύ':   'y',
	'Ώ':   'o',
	'ΐ':   'i',
	'Α':   'a',
	'Β':   'b',
	'Γ':   't',
	'Δ':   'd',
	'Ε':   'e',
	'Ζ':   'z',
	'Η':   'h',
	'Θ':   'o',
	'Ι':   'i',
	'Κ':   'k',
	'Λ':   'a',
	'Μ':   'm',
	'Ν':   'n',
	'Ξ':   'x',
	'Ο':   'o',
	'Π':   'p',
	'Ρ':   'p',
	0x3a2: '.',
	'Σ':   'e',
	'Τ':   't',
	'Υ':   'y',
	'Φ':   'o',
	'Χ':   'x',
	'Ψ':   'w',
	'Ω':   'o',
	'Ϊ':   'i',
	'Ϋ':   'y',
	'ά':   'a',
	'έ':   'e',
	'ή':   'n',
	'ί':   'i',
	'ΰ':   'u',
	'α':   'a',
	'β':   'b',
	'γ':   'y',
	'δ':   'o',
	'ε':   'e',
	'ζ':   's',
	'η':   'n',
	'θ':   'o',
	'ι':   'i',
	'κ':   'k',
	'λ':   'a',
	'μ':   'u',
	'ν':   'v',
	'ξ':   'e',
	'ο':   'o',
	'π':   'p',
	'ρ':   'p',
	'ς':   's',
	'σ':   'o',
	'τ':   't',
	'υ':   'u',
	'φ':   'o',
	'χ':   'x',
	'ψ':   'w',
	'ω':   'w',
	'ϊ':   'i',
	'ϋ':   'u',
	'ό':   'o',
	'ύ':   'u',
	'ώ':   'w',
	'Ϗ':   'k',
	'ϐ':   'b',
	'ϑ':   'o',
	'ϒ':   'y',
	'ϓ':   'y',
	'ϔ':   'y',
	'ϕ':   'o',
	'ϖ':   'w',
	'ϗ':   'x',
	'Ϙ':   'o',
	'ϙ':   'o',
	'Ϛ':   's',
	'ϛ':   's',
	'Ϝ':   'f',
	'ϝ':   'f',
	'Ϟ':   'z',
	'ϟ':   's',
	'Ϡ':   'e',
	'ϡ':   'e',
	'Ϣ':   'w',
	'ϣ':   'w',
	'Ϥ':   'y',
	'ϥ':   'y',
	'Ϧ':   'h',
	'ϧ':   'h',
	'Ϩ':   'z',
	'ϩ':   'z',
	'Ϫ':   'a',
	'ϫ':   'a',
	'Ϭ':   'b',
	'ϭ':   'b',
	'Ϯ':   't',
	'ϯ':   't',
	'ϰ':   'x',
	'ϱ':   'p',
	'ϲ':   'c',
	'ϳ':   'j',
	'ϴ':   'o',
	'ϵ':   'e',
	'϶':   'e',
	'Ϸ':   'p',
	'ϸ':   'p',
	'Ϲ':   'c',
	'Ϻ':   'm',
	'ϻ':   'm',
	'ϼ':   'p',
	'Ͻ':   'c',
	'Ͼ':   'c',
	'Ͽ':   'c',
	'Ѐ':   'e',
	'Ё':   'e',
	'Ђ':   'h',
	'Ѓ':   't',
	'Є':   'e',
	'Ѕ':   's',
	'І':   'i',
	'Ї':   'i',
	'Ј':   'j',
	'Љ':   'b',
	'Њ':   'b',
	'Ћ':   'h',
	'Ќ':   'k',
	'Ѝ':   'n',
	'Ў':   'y',
	'Џ':   'u',
	'А':   'a',
	'Б':   'b',
	'В':   'b',
	'Г':   't',
	'Д':   'd',
	'Е':   'e',
	'Ж':   'k',
	'З':   'b',
	'И':   'n',
	'Й':   'n',
	'К':   'k',
	'Л':   'a',
	'М':   'm',
	'Н':   'h',
	'О':   'o',
	'П':   'p',
	'Р':   'p',
	'С':   'c',
	'Т':   't',
	'У':   'y',
	'Ф':   'p',
	'Х':   'x',
	'Ц':   'u',
	'Ч':   'y',
	'Ш':   'w',
	'Щ':   'w',
	'Ъ':   'b',
	'Ы':   'b',
	'Ь':   'b',
	'Э':   'e',
	'Ю':   'o',
	'Я':   'r',
	'а':   'a',
	'б':   'b',
	'в':   'b',
	'г':   't',
	'д':   'd',
	'е':   'e',
	'ж':   'k',
	'з':   'b',
	'и':   'n',
	'й':   'n',
	'к':   'k',
	'л':   'a',
	'м':   'm',
	'н':   'h',
	'о':   'o',
	'п':   'p',
	'р':   'p',
	'с':   'c',
	'т':   't',
	'у':   'y',
	'ф':   'p',
	'х':   'x',
	'ц':   'u',
	'ч':   'y',
	'ш':   'w',
	'щ':   'w',
	'ъ':   'b',
	'ы':   'b',
	'ь':   'b',
	'э':   'e',
	'ю':   'o',
	'я':   'r',
	'ѐ':   'e',
	'ё':   'e',
	'ђ':   'h',
	'ѓ':   't',
	'є':   'e',
	'ѕ':   's',
	'і':   'i',
	'ї':   'i',
	'ј':   'j',
	'љ':   'b',
	'њ':   'b',
	'ћ':   'h',
	'ќ':   'k',
	'ѝ':   'n',
	'ў':   'y',
	'џ':   'u',
	'Ѡ':   'w',
	'ѡ':   'w',
	'Ѣ':   'b',
	'ѣ':   'b',
	'Ѥ':   'e',
	'ѥ':   'e',
	'Ѧ':   'a',
	'ѧ':   'a',
	'Ѩ':   'a',
	'ѩ':   'a',
	'Ѫ':   'e',
	'ѫ':   'e',
	'Ѭ':   'e',
	'ѭ':   'e',
	'Ѯ':   'b',
	'ѯ':   'b',
	'Ѱ':   'w',
	'ѱ':   'w',
	'Ѳ':   'o',
	'ѳ':   'o',
	'Ѵ':   'v',
	'ѵ':   'v',
	'Ѷ':   'v',
	'ѷ':   'v',
	'Ѹ':   'y',
	'ѹ':   'y',
	'Ѻ':   'o',
	'ѻ':   'o',
	'Ѽ':   'w',
	'ѽ':   'w',
	'Ѿ':   'w',
	'ѿ':   'w',
	'Ҁ':   'c',
	'ҁ':   'c',
	'҂':   '.',
	0x483: '.',
	0x484: '.',
	0x485: '.',
	0x486: '.',
	0x487: '.',
	0x488: '.',
	0x489: '.',
	'Ҋ':   'n',
	'ҋ':   'n',
	'Ҍ':   'b',
	'ҍ':   'b',
	'Ҏ':   'p',
	'ҏ':   'p',
	'Ґ':   't',
	'ґ':   't',
	'Ғ':   'f',
	'ғ':   'f',
	'Ҕ':   'h',
	'ҕ':   'h',
	'Җ':   'k',
	'җ':   'k',
	'Ҙ':   'b',
	'ҙ':   'b',
	'Қ':   'k',
	'қ':   'k',
	'Ҝ':   'k',
	'ҝ':   'k',
	'Ҟ':   'k',
	'ҟ':   'k',
	'Ҡ':   'k',
	'ҡ':   'k',
	'Ң':   'h',
	'ң':   'h',
	'Ҥ':   'h',
	'ҥ':   'h',
	'Ҧ':   'm',
	'ҧ':   'm',
	'Ҩ':   'a',
	'ҩ':   'a',
	'Ҫ':   'c',
	'ҫ':   'c',
	'Ҭ':   't',
	'ҭ':   't',
	'Ү':   'y',
	'ү':   'y',
	'Ұ':   'y',
	'ұ':   'y',
	'Ҳ':   'x',
	'ҳ':   'x',
	'Ҵ':   'u',
	'ҵ':   'u',
	'Ҷ':   'y',
	'ҷ':   'y',
	'Ҹ':   'y',
	'ҹ':   'y',
	'Һ':   'h',
	'һ':   'h',
	'Ҽ':   'e',
	'ҽ':   'e',
	'Ҿ':   'e',
	'ҿ':   'e',
	'Ӏ':   'i',
	'Ӂ':   'k',
	'ӂ':   'k',
	'Ӄ':   'k',
	'ӄ':   'k',
	'Ӆ':   'a',
	'ӆ':   'a',
	'Ӈ':   'h',
	'ӈ':   'h',
	'Ӊ':   'h',
	'ӊ':   'h',
	'Ӌ':   'y',
	'ӌ':   'y',
	'Ӎ':   'm',
	'ӎ':   'm',
	'ӏ':   'i',
	'Ӑ':   'a',
	'ӑ':   'a',
	'Ӓ':   'a',
	'ӓ':   'a',
	'Ӕ':   'e',
	'ӕ':   'e',
	'Ӗ':   'e',
	'ӗ':   'e',
	'Ә':   'e',
	'ә':   'e',
	'Ӛ':   'e',
	'ӛ':   'e',
	'Ӝ':   'k',
	'ӝ':   'k',
	'Ӟ':   'b',
	'ӟ':   'b',
	'Ӡ':   'z',
	'ӡ':   'z',
	'Ӣ':   'n',
	'ӣ':   'n',
	'Ӥ':   'n',
	'ӥ':   'n',
	'Ӧ':   'o',
	'ӧ':   'o',
	'Ө':   'o',
	'ө':   'o',
	'Ӫ':   'o',
	'ӫ':   'o',
	'Ӭ':   'e',
	'ӭ':   'e',
	'Ӯ':   'y',
	'ӯ':   'y',
	'Ӱ':   'y',
	'ӱ':   'y',
	'Ӳ':   'y',
	'ӳ':   'y',
	'Ӵ':   'y',
	'ӵ':   'y',
	'Ӷ':   't',
	'ӷ':   't',
	'Ӹ':   'b',
	'ӹ':   'b',
	'Ӻ':   'f',
	'ӻ':   'f',
	'Ӽ':   'x',
	'ӽ':   'x',
	'Ӿ':   'x',
	'ӿ':   'x',
	'Ԁ':   'd',
	'ԁ':   'd',
	'Ԃ':   'd',
	'ԃ':   'd',
	'Ԅ':   'u',
	'ԅ':   'u',
	'Ԇ':   'r',
	'ԇ':   'r',
	'Ԉ':   'n',
	'ԉ':   'n',
	'Ԋ':   'h',
	'ԋ':   'h',
	'Ԍ':   'g',
	'ԍ':   'g',
	'Ԏ':   't',
	'ԏ':   't',
	'Ԑ':   'e',
	'ԑ':   'e',
	'Ԓ':   'a',
	'ԓ':   'a',
	'Ԕ':   'x',
	'ԕ':   'x',
	'Ԗ':   'p',
	'ԗ':   'p',
	'Ԙ':   'e',
	'ԙ':   'e',
	'Ԛ':   'q',
	'ԛ':   'q',
	'Ԝ':   'w',
	'ԝ':   'w',
	'Ԟ':   'k',
	'ԟ':   'k',
	'Ԡ':   'm',
	'ԡ':   'm',
	'Ԣ':   'h',
	'ԣ':   'h',
	'Ԥ':   'p',
	'ԥ':   'p',
	'Ԧ':   'h',
	'ԧ':   'h',
	'Ԩ':   'h',
	'ԩ':   'h',
	'Ԫ':   'k',
	'ԫ':   'k',
	'Ԭ':   'a',
	'ԭ':   'a',
	'Ԯ':   'a',
	'ԯ':   'a',
}

var kTranslitB = [kVietnameseEnd1 - kVietnameseStart]rune{
	'Ḁ' - kVietnameseStart: 'a',
	'ḁ' - kVietnameseStart: 'a',
	'Ḃ' - kVietnameseStart: 'b',
	'ḃ' - kVietnameseStart: 'b',
	'Ḅ' - kVietnameseStart: 'b',
	'ḅ' - kVietnameseStart: 'b',
	'Ḇ' - kVietnameseStart: 'b',
	'ḇ' - kVietnameseStart: 'b',
	'Ḉ' - kVietnameseStart: 'c',
	'ḉ' - kVietnameseStart: 'c',
	'Ḋ' - kVietnameseStart: 'd',
	'ḋ' - kVietnameseStart: 'd',
	'Ḍ' - kVietnameseStart: 'd',
	'ḍ' - kVietnameseStart: 'd',
	'Ḏ' - kVietnameseStart: 'd',
	'ḏ' - kVietnameseStart: 'd',
	'Ḑ' - kVietnameseStart: 'd',
	'ḑ' - kVietnameseStart: 'd',
	'Ḓ' - kVietnameseStart: 'd',
	'ḓ' - kVietnameseStart: 'd',
	'Ḕ' - kVietnameseStart: 'e',
	'ḕ' - kVietnameseStart: 'e',
	'Ḗ' - kVietnameseStart: 'e',
	'ḗ' - kVietnameseStart: 'e',
	'Ḙ' - kVietnameseStart: 'e',
	'ḙ' - kVietnameseStart: 'e',
	'Ḛ' - kVietnameseStart: 'e',
	'ḛ' - kVietnameseStart: 'e',
	'Ḝ' - kVietnameseStart: 'e',
	'ḝ' - kVietnameseStart: 'e',
	'Ḟ' - kVietnameseStart: 'f',
	'ḟ' - kVietnameseStart: 'f',
	'Ḡ' - kVietnameseStart: 'g',
	'ḡ' - kVietnameseStart: 'g',
	'Ḣ' - kVietnameseStart: 'h',
	'ḣ' - kVietnameseStart: 'h',
	'Ḥ' - kVietnameseStart: 'h',
	'ḥ' - kVietnameseStart: 'h',
	'Ḧ' - kVietnameseStart: 'h',
	'ḧ' - kVietnameseStart: 'h',
	'Ḩ' - kVietnameseStart: 'h',
	'ḩ' - kVietnameseStart: 'h',
	'Ḫ' - kVietnameseStart: 'h',
	'ḫ' - kVietnameseStart: 'h',
	'Ḭ' - kVietnameseStart: 'i',
	'ḭ' - kVietnameseStart: 'i',
	'Ḯ' - kVietnameseStart: 'i',
	'ḯ' - kVietnameseStart: 'i',
	'Ḱ' - kVietnameseStart: 'k',
	'ḱ' - kVietnameseStart: 'k',
	'Ḳ' - kVietnameseStart: 'k',
	'ḳ' - kVietnameseStart: 'k',
	'Ḵ' - kVietnameseStart: 'k',
	'ḵ' - kVietnameseStart: 'k',
	'Ḷ' - kVietnameseStart: 'l',
	'ḷ' - kVietnameseStart: 'l',
	'Ḹ' - kVietnameseStart: 'l',
	'ḹ' - kVietnameseStart: 'l',
	'Ḻ' - kVietnameseStart: 'l',
	'ḻ' - kVietnameseStart: 'l',
	'Ḽ' - kVietnameseStart: 'l',
	'ḽ' - kVietnameseStart: 'l',
	'Ḿ' - kVietnameseStart: 'm',
	'ḿ' - kVietnameseStart: 'm',
	'Ṁ' - kVietnameseStart: 'm',
	'ṁ' - kVietnameseStart: 'm',
	'Ṃ' - kVietnameseStart: 'm',
	'ṃ' - kVietnameseStart: 'm',
	'Ṅ' - kVietnameseStart: 'n',
	'ṅ' - kVietnameseStart: 'n',
	'Ṇ' - kVietnameseStart: 'n',
	'ṇ' - kVietnameseStart: 'n',
	'Ṉ' - kVietnameseStart: 'n',
	'ṉ' - kVietnameseStart: 'n',
	'Ṋ' - kVietnameseStart: 'n',
	'ṋ' - kVietnameseStart: 'n',
	'Ṍ' - kVietnameseStart: 'o',
	'ṍ' - kVietnameseStart: 'o',
	'Ṏ' - kVietnameseStart: 'o',
	'ṏ' - kVietnameseStart: 'o',
	'Ṑ' - kVietnameseStart: 'o',
	'ṑ' - kVietnameseStart: 'o',
	'Ṓ' - kVietnameseStart: 'o',
	'ṓ' - kVietnameseStart: 'o',
	'Ṕ' - kVietnameseStart: 'p',
	'ṕ' - kVietnameseStart: 'p',
	'Ṗ' - kVietnameseStart: 'p',
	'ṗ' - kVietnameseStart: 'p',
	'Ṙ' - kVietnameseStart: 'r',
	'ṙ' - kVietnameseStart: 'r',
	'Ṛ' - kVietnameseStart: 'r',
	'ṛ' - kVietnameseStart: 'r',
	'Ṝ' - kVietnameseStart: 'r',
	'ṝ' - kVietnameseStart: 'r',
	'Ṟ' - kVietnameseStart: 'r',
	'ṟ' - kVietnameseStart: 'r',
	'Ṡ' - kVietnameseStart: 's',
	'ṡ' - kVietnameseStart: 's',
	'Ṣ' - kVietnameseStart: 's',
	'ṣ' - kVietnameseStart: 's',
	'Ṥ' - kVietnameseStart: 's',
	'ṥ' - kVietnameseStart: 's',
	'Ṧ' - kVietnameseStart: 's',
	'ṧ' - kVietnameseStart: 's',
	'Ṩ' - kVietnameseStart: 's',
	'ṩ' - kVietnameseStart: 's',
	'Ṫ' - kVietnameseStart: 't',
	'ṫ' - kVietnameseStart: 't',
	'Ṭ' - kVietnameseStart: 't',
	'ṭ' - kVietnameseStart: 't',
	'Ṯ' - kVietnameseStart: 't',
	'ṯ' - kVietnameseStart: 't',
	'Ṱ' - kVietnameseStart: 't',
	'ṱ' - kVietnameseStart: 't',
	'Ṳ' - kVietnameseStart: 'u',
	'ṳ' - kVietnameseStart: 'u',
	'Ṵ' - kVietnameseStart: 'u',
	'ṵ' - kVietnameseStart: 'u',
	'Ṷ' - kVietnameseStart: 'u',
	'ṷ' - kVietnameseStart: 'u',
	'Ṹ' - kVietnameseStart: 'u',
	'ṹ' - kVietnameseStart: 'u',
	'Ṻ' - kVietnameseStart: 'u',
	'ṻ' - kVietnameseStart: 'u',
	'Ṽ' - kVietnameseStart: 'v',
	'ṽ' - kVietnameseStart: 'v',
	'Ṿ' - kVietnameseStart: 'v',
	'ṿ' - kVietnameseStart: 'v',
	'Ẁ' - kVietnameseStart: 'w',
	'ẁ' - kVietnameseStart: 'w',
	'Ẃ' - kVietnameseStart: 'w',
	'ẃ' - kVietnameseStart: 'w',
	'Ẅ' - kVietnameseStart: 'w',
	'ẅ' - kVietnameseStart: 'w',
	'Ẇ' - kVietnameseStart: 'w',
	'ẇ' - kVietnameseStart: 'w',
	'Ẉ' - kVietnameseStart: 'w',
	'ẉ' - kVietnameseStart: 'w',
	'Ẋ' - kVietnameseStart: 'x',
	'ẋ' - kVietnameseStart: 'x',
	'Ẍ' - kVietnameseStart: 'x',
	'ẍ' - kVietnameseStart: 'x',
	'Ẏ' - kVietnameseStart: 'y',
	'ẏ' - kVietnameseStart: 'y',
	'Ẑ' - kVietnameseStart: 'z',
	'ẑ' - kVietnameseStart: 'z',
	'Ẓ' - kVietnameseStart: 'z',
	'ẓ' - kVietnameseStart: 'z',
	'Ẕ' - kVietnameseStart: 'z',
	'ẕ' - kVietnameseStart: 'z',
	'ẖ' - kVietnameseStart: 'h',
	'ẗ' - kVietnameseStart: 't',
	'ẘ' - kVietnameseStart: 'w',
	'ẙ' - kVietnameseStart: 'y',
	'ẚ' - kVietnameseStart: 'a',
	'ẛ' - kVietnameseStart: 's',
	'ẜ' - kVietnameseStart: 'f',
	'ẝ' - kVietnameseStart: 'f',
	'ẞ' - kVietnameseStart: 'b',
	'ẟ' - kVietnameseStart: 'o',
	'Ạ' - kVietnameseStart: 'a',
	'ạ' - kVietnameseStart: 'a',
	'Ả' - kVietnameseStart: 'a',
	'ả' - kVietnameseStart: 'a',
	'Ấ' - kVietnameseStart: 'a',
	'ấ' - kVietnameseStart: 'a',
	'Ầ' - kVietnameseStart: 'a',
	'ầ' - kVietnameseStart: 'a',
	'Ẩ' - kVietnameseStart: 'a',
	'ẩ' - kVietnameseStart: 'a',
	'Ẫ' - kVietnameseStart: 'a',
	'ẫ' - kVietnameseStart: 'a',
	'Ậ' - kVietnameseStart: 'a',
	'ậ' - kVietnameseStart: 'a',
	'Ắ' - kVietnameseStart: 'a',
	'ắ' - kVietnameseStart: 'a',
	'Ằ' - kVietnameseStart: 'a',
	'ằ' - kVietnameseStart: 'a',
	'Ẳ' - kVietnameseStart: 'a',
	'ẳ' - kVietnameseStart: 'a',
	'Ẵ' - kVietnameseStart: 'a',
	'ẵ' - kVietnameseStart: 'a',
	'Ặ' - kVietnameseStart: 'a',
	'ặ' - kVietnameseStart: 'a',
	'Ẹ' - kVietnameseStart: 'e',
	'ẹ' - kVietnameseStart: 'e',
	'Ẻ' - kVietnameseStart: 'e',
	'ẻ' - kVietnameseStart: 'e',
	'Ẽ' - kVietnameseStart: 'e',
	'ẽ' - kVietnameseStart: 'e',
	'Ế' - kVietnameseStart: 'e',
	'ế' - kVietnameseStart: 'e',
	'Ề' - kVietnameseStart: 'e',
	'ề' - kVietnameseStart: 'e',
	'Ể' - kVietnameseStart: 'e',
	'ể' - kVietnameseStart: 'e',
	'Ễ' - kVietnameseStart: 'e',
	'ễ' - kVietnameseStart: 'e',
	'Ệ' - kVietnameseStart: 'e',
	'ệ' - kVietnameseStart: 'e',
	'Ỉ' - kVietnameseStart: 'i',
	'ỉ' - kVietnameseStart: 'i',
	'Ị' - kVietnameseStart: 'i',
	'ị' - kVietnameseStart: 'i',
	'Ọ' - kVietnameseStart: 'o',
	'ọ' - kVietnameseStart: 'o',
	'Ỏ' - kVietnameseStart: 'o',
	'ỏ' - kVietnameseStart: 'o',
	'Ố' - kVietnameseStart: 'o',
	'ố' - kVietnameseStart: 'o',
	'Ồ' - kVietnameseStart: 'o',
	'ồ' - kVietnameseStart: 'o',
	'Ổ' - kVietnameseStart: 'o',
	'ổ' - kVietnameseStart: 'o',
	'Ỗ' - kVietnameseStart: 'o',
	'ỗ' - kVietnameseStart: 'o',
	'Ộ' - kVietnameseStart: 'o',
	'ộ' - kVietnameseStart: 'o',
	'Ớ' - kVietnameseStart: 'o',
	'ớ' - kVietnameseStart: 'o',
	'Ờ' - kVietnameseStart: 'o',
	'ờ' - kVietnameseStart: 'o',
	'Ở' - kVietnameseStart: 'o',
	'ở' - kVietnameseStart: 'o',
	'Ỡ' - kVietnameseStart: 'o',
	'ỡ' - kVietnameseStart: 'o',
	'Ợ' - kVietnameseStart: 'o',
	'ợ' - kVietnameseStart: 'o',
	'Ụ' - kVietnameseStart: 'u',
	'ụ' - kVietnameseStart: 'u',
	'Ủ' - kVietnameseStart: 'u',
	'ủ' - kVietnameseStart: 'u',
	'Ứ' - kVietnameseStart: 'u',
	'ứ' - kVietnameseStart: 'u',
	'Ừ' - kVietnameseStart: 'u',
	'ừ' - kVietnameseStart: 'u',
	'Ử' - kVietnameseStart: 'u',
	'ử' - kVietnameseStart: 'u',
	'Ữ' - kVietnameseStart: 'u',
	'ữ' - kVietnameseStart: 'u',
	'Ự' - kVietnameseStart: 'u',
	'ự' - kVietnameseStart: 'u',
	'Ỳ' - kVietnameseStart: 'y',
	'ỳ' - kVietnameseStart: 'y',
	'Ỵ' - kVietnameseStart: 'y',
	'ỵ' - kVietnameseStart: 'y',
	'Ỷ' - kVietnameseStart: 'y',
	'ỷ' - kVietnameseStart: 'y',
	'Ỹ' - kVietnameseStart: 'y',
	'ỹ' - kVietnameseStart: 'y',
	'Ỻ' - kVietnameseStart: 'l',
	'ỻ' - kVietnameseStart: 'l',
	'Ỽ' - kVietnameseStart: 'b',
	'ỽ' - kVietnameseStart: 'b',
	'Ỿ' - kVietnameseStart: 'y',
	'ỿ' - kVietnameseStart: 'y',
}
