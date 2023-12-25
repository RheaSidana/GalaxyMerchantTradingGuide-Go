package romanNumerals

import "fmt"

func BuildRomanNumerals() {
	createRomanNumeralsEquivalent()
	createRomanNumeralsNotRepeated()
}

func ConvertToInt(command []string, invalidFormat string) int {
	romanNumeral := command[1]

	prev := 0
	res := 0
	repeatedCount := 0
	for _, numeral := range romanNumeral {
		curr := int(RomanNumeralsEquivalent[numeral])
		if curr == prev{
			repeatedCount ++
			if repeatedCount > 2 || RomanNumeralsNotRepeated[numeral]{
				fmt.Println(invalidFormat)
				return 0
			}
		}

		if prev < curr {
			if (prev == 5 || prev == 50 || prev == 500) {
				fmt.Println(invalidFormat)
				return 0
			}
			res += curr - (2 * prev)
		} else {
			res += curr
		}

		prev = curr
	}

	fmt.Println(command[1] , " is ", res)
	return res
}

