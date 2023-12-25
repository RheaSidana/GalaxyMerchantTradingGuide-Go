package intergalactic

import (
	// "fmt"
	"strconv"
)

func InitIntergalactic() {
	createQues()
}

func CreateSymbols(command []string) {
	symbols[command[0]] = command[2]
	// fmt.Println("len", len(command[0]))
	// fmt.Println("Symbol["+command[0]+"]:", symbols[command[0]])
}

func CreateMetals(command []string, convertToInt func([]string, string)int, REQUEST_INVALID_FORMAT string) {
	// symbols[command[0]] = command[2]
	// fmt.Println("len", len(command[0]))
	// fmt.Println("Metal: ", command)

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

	// fmt.Println("Metals["+metal+"]", metals[metal])
}

func calculateOneMetalEquals(
	command []string,
	response string,
	convertToInt func([]string, string)int,
	REQUEST_INVALID_FORMAT string,
) float64 {

	convertCommand := []string{"Convert", response}
	denom := convertToInt(convertCommand, REQUEST_INVALID_FORMAT)

	credits,_ := strconv.ParseInt(command[len(command)-2], 10, 64)
	metalValue := float64(credits) / float64((denom))
	return metalValue
}
