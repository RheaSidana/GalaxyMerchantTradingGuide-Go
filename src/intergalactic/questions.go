package intergalactic

import (
	"fmt"
	"strings"
)

func HowMuchIsQues(
	command []string,
	convertToInt func([]string, string) int,
	REQUEST_INVALID_FORMAT string,
) {
	// fmt.Println("in how much is")
	responseString := ""
	romanString := ""
	for i := 3; i < len(command)-1; i++ {
		// fmt.Println("comand :" + command[i])
		responseString += command[i] + " "
		if val, ok := symbols[command[i]]; ok {
			romanString += val
		}
	}

	conversionRate := conversionRateFromRoman(
		romanString,
		convertToInt,
		REQUEST_INVALID_FORMAT,
	)

	fmt.Println(responseString+"is", conversionRate)
}

func HowManyCreditsIsQues(
	command []string,
	convertToInt func([]string, string) int,
	REQUEST_INVALID_FORMAT string,
) {
	// fmt.Println("in how much is")
	responseString := ""
	romanString := ""
	metalRate := 0.0
	for i := 4; i < len(command)-1; i++ {
		// fmt.Println("comand :" + command[i])
		responseString += command[i] + " "
		if val, ok := symbols[command[i]]; ok {
			romanString += val
		} else {
			metalRate = metals[command[i]]
		}
	}

	conversionRate := conversionRateFromRoman(
		romanString,
		convertToInt,
		REQUEST_INVALID_FORMAT,
	)

	metalAmount := conversionRate * metalRate

	if metalAmount != 0 {
		fmt.Println(responseString+"is", metalAmount, "Credits")
	}
}

var DoesHasMap = map[string]string{
	"has more Credits than": "has less Credits than",
	"has less Credits than": "has more Credits than",
}

func DoesHasMoreOrLessQues(
	command []string,
	convertToInt func([]string, string) int,
	REQUEST_INVALID_FORMAT string,
) {
	// fmt.Println("in how much is")
	// responseString := ""
	firstElement, secondElement,
		romanStringOfFirstElement,
		romanStringOfSecondElement,
		hasString := extractComparedStringDetails(command,
		"has",
		"!",
	)

	words := strings.Fields(firstElement)
	metalRateOfFirstElement := metals[words[len(words)-1]]
	words = strings.Fields(secondElement)
	metalRateOfSecondElement := metals[words[len(words)-1]]

	conversionRateOfFirstElement := conversionRateFromRoman(
		romanStringOfFirstElement,
		convertToInt,
		REQUEST_INVALID_FORMAT,
	)
	conversionRateOfSecondElement := conversionRateFromRoman(
		romanStringOfSecondElement,
		convertToInt,
		REQUEST_INVALID_FORMAT,
	)

	metalAmountFirstElement := conversionRateOfFirstElement * metalRateOfFirstElement
	metalAmountSecondElement := conversionRateOfSecondElement * metalRateOfSecondElement

	if strings.Contains(hasString, "more") &&
		metalAmountFirstElement < metalAmountSecondElement {
		hasString = DoesHasMap[hasString]
	} else if strings.Contains(hasString, "less") &&
		metalAmountFirstElement > metalAmountSecondElement {
		hasString = DoesHasMap[hasString]
	}

	fmt.Println(firstElement + hasString + secondElement)
}

func extractComparedStringDetails(command []string, word string, alt string) (
	string,
	string,
	string,
	string,
	string,
) {
	firstElement := ""
	secondElement := ""
	romanStringOfFirstElement := ""
	romanStringOfSecondElement := ""
	hasString := ""
	firstElementFound := false
	thanFound := false
	for i := 1; i < len(command)-2; i++ {
		// fmt.Println("comand :" + command[i])
		if command[i] == word || command[i] == alt {
			firstElementFound = true
		}
		if command[i] == "than" {
			hasString += command[i]
			thanFound = true
		}

		if !firstElementFound {

			firstElement += command[i] + " "
			if val, ok := symbols[command[i]]; ok {
				romanStringOfFirstElement += val
			}

		} else if firstElementFound && !thanFound {

			hasString += command[i] + " "
			// fmt.Println(hasString)

		} else if firstElementFound && thanFound {

			secondElement += " " + command[i+1]
			if val, ok := symbols[command[i+1]]; ok {
				romanStringOfSecondElement += val
			}

		}

	}

	return firstElement,
		secondElement,
		romanStringOfFirstElement,
		romanStringOfSecondElement,
		hasString
}

var IsMap = map[string]string{
	"larger than":  "smaller than",
	"smaller than": "larger than",
}

func IsQues(
	command []string,
	convertToInt func([]string, string) int,
	REQUEST_INVALID_FORMAT string,
) {
	if len(command[0]) > 2 {
		command = append(
			[]string{
				"Is",
				strings.Replace(command[0], "Is", "", 1),
			},
			command[1:]...,
		)
	}
	firstElement, secondElement,
		romanStringOfFirstElement,
		romanStringOfSecondElement,
		hasString := extractComparedStringDetails(command,
		"larger", "smaller")

	conversionRateOfFirstElement := conversionRateFromRoman(
		romanStringOfFirstElement,
		convertToInt,
		REQUEST_INVALID_FORMAT,
	)
	conversionRateOfSecondElement := conversionRateFromRoman(
		romanStringOfSecondElement,
		convertToInt,
		REQUEST_INVALID_FORMAT,
	)

	// metalAmountFirstElement := conversionRateOfFirstElement * metalRateOfFirstElement
	// metalAmountSecondElement := conversionRateOfSecondElement * metalRateOfSecondElement

	if strings.Contains(hasString, "larger") &&
		conversionRateOfFirstElement < conversionRateOfSecondElement {
		hasString = IsMap[hasString]
	} else if strings.Contains(hasString, "smaller") &&
		conversionRateOfFirstElement > conversionRateOfSecondElement {
		hasString = IsMap[hasString]
	}

	fmt.Println(firstElement + hasString + secondElement)
}
