package romanNumerals

var RomanNumeralsEquivalent = make(map[rune]int64)
var RomanNumeralsNotRepeated = make(map[rune]bool)

func createRomanNumeralsEquivalent() {
	RomanNumeralsEquivalent['I'] = 1
	RomanNumeralsEquivalent['V'] = 5
	RomanNumeralsEquivalent['X'] = 10
	RomanNumeralsEquivalent['L'] = 50
	RomanNumeralsEquivalent['C'] = 100
	RomanNumeralsEquivalent['D'] = 500
	RomanNumeralsEquivalent['M'] = 1000
}

func createRomanNumeralsNotRepeated() {
	RomanNumeralsNotRepeated['V'] = true
	RomanNumeralsNotRepeated['L'] = true
	RomanNumeralsNotRepeated['D'] = true
}
