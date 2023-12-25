package commands

import romanNumerals "galaxy_merchant_trading_guide/src/romanNumerals"

const REQUEST_INVALID_FORMAT = "Request number is in invalid format"
const REQUEST_INVALID_QUESTION = "I have no idea what you are talking about"

var commandsMap = make(map[string]func([]string) int)

const COMMAND_FOUND = 1

func CreateCommands() {
	commandsMap["Convert"] = func(command []string) int {
		romanNumerals.ConvertToInt(command, REQUEST_INVALID_FORMAT)
		return COMMAND_FOUND
	}
}