package intergalactic

import (
	// "fmt"
	"strconv"
)

func CreateSymbols(command []string) {
	symbols[command[0]] = command[2]
}

func CreateMetals(
	command []string,
	convertToInt func([]string, string) int,
	REQUEST_INVALID_FORMAT string,
) {
	res := ""
	metal := ""
	for _, word := range command {
		if val, ok := symbols[word]; ok {
			res += val
		} else {
			metal = word
			break
		}
	}

	metalValue := calculateOneMetalEquals(
		command,
		res,
		convertToInt,
		REQUEST_INVALID_FORMAT,
	)

	metals[metal] = metalValue
}

func calculateOneMetalEquals(
	command []string,
	response string,
	convertToInt func([]string, string) int,
	REQUEST_INVALID_FORMAT string,
) float64 {

	denom := conversionRateFromRoman(response, convertToInt, REQUEST_INVALID_FORMAT)

	credits, _ := strconv.ParseInt(command[len(command)-2], 10, 64)
	metalValue := float64(credits) / float64((denom))
	return metalValue
}

func conversionRateFromRoman(
	response string,
	convertToInt func([]string, string) int,
	REQUEST_INVALID_FORMAT string,
) float64 {
	convertCommand := []string{"Convert", response}
	denom := float64(convertToInt(convertCommand, REQUEST_INVALID_FORMAT))

	return denom
}
