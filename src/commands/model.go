package commands

import (
	"fmt"
	intergalactic "galaxy_merchant_trading_guide/src/intergalactic"
	romanNumerals "galaxy_merchant_trading_guide/src/romanNumerals"
)

const REQUEST_INVALID_FORMAT = "Request number is in invalid format"
const REQUEST_INVALID_QUESTION = "I have no idea what you are talking about"

var commandsMap = make(map[string]func([]string) int)
var intergalacticMap = make(map[string]func([]string) int)

const COMMAND_FOUND = 1
const SYMBOL_FOUND = 1

func CreateCommands() {
	commandsMap["Convert"] = func(command []string) int {
		conv := romanNumerals.ConvertToInt(command, REQUEST_INVALID_FORMAT)
		if conv != 0 {
			fmt.Println(command[1], " is ", conv)
		}
		return COMMAND_FOUND
	}

	commandsMap["how much is"] = func(command []string) int {
		intergalactic.HowMuchIsQues(
			command,
			romanNumerals.ConvertToInt,
			REQUEST_INVALID_FORMAT,
		)
		return COMMAND_FOUND
	}

	commandsMap["how many Credits is"] = func(command []string) int {
		intergalactic.HowManyCreditsIsQues(
			command,
			romanNumerals.ConvertToInt,
			REQUEST_INVALID_FORMAT,
		)
		return COMMAND_FOUND
	}

	commandsMap["Does"] = func(command []string) int {
		intergalactic.DoesHasMoreOrLessQues(
			command,
			romanNumerals.ConvertToInt,
			REQUEST_INVALID_FORMAT,
		)
		return COMMAND_FOUND
	}

	commandsMap["Is"] = func(command []string) int {
		intergalactic.IsQues(
			command,
			romanNumerals.ConvertToInt,
			REQUEST_INVALID_FORMAT,
		)
		return COMMAND_FOUND
	}
}

func CreateSymbols() {
	intergalacticMap["is"] = func(symbols []string) int {
		intergalactic.CreateSymbols(symbols)
		return SYMBOL_FOUND
	}

	intergalacticMap["Credits"] = func(symbols []string) int {
		intergalactic.CreateMetals(symbols, romanNumerals.ConvertToInt, REQUEST_INVALID_FORMAT)
		return SYMBOL_FOUND
	}
}
