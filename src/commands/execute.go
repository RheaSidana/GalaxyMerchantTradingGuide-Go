package commands

import (
	"fmt"
	"strings"
)

func Execute(batchCommands [][]string) {
	for _, commandPerLine := range batchCommands {
		if isBlankLine(commandPerLine) {
			continue
		}

		// fmt.Print("\nCommand: ")
		// fmt.Println(commandPerLine)
		err := ProcessCommandIfAvailable(commandPerLine)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func isBlankLine(line []string) bool {
	blankLineLength := 0
	return len(line) == blankLineLength
}

func ProcessCommandIfAvailable(command []string) error {
	// fmt.Println("\nCommand : ", command)
	result := commandsMap[command[0]]

	if !isValidCommand(result) {
		symbol := intergalacticMap[command[1]]
		credits := intergalacticMap[command[len(command)-1]]
		hasFoundIntergalacticSymbol := hasFoundIntergalactic(symbol)
		hasFoundIntergalacticCredit := hasFoundIntergalactic(credits)

		isValidQues, res := isValidQues(command)
		// fmt.Println("isValidQues: ", isValidQues)
		// fmt.Println("res: ", res)
		if !hasFoundIntergalacticSymbol &&
			!hasFoundIntergalacticCredit &&
			!isValidQues {

			fmt.Println(REQUEST_INVALID_QUESTION)
			return nil
		}

		if hasFoundIntergalacticSymbol {
			result = symbol
		} else if hasFoundIntergalacticCredit {
			result = credits
		} else {
			result = res
		}
	}

	return executeCommand(result, command)
}

func isValidCommand(result func([]string) int) bool {
	return result != nil
}

func executeCommand(result func([]string) int, command []string) error {
	result(command)
	return nil
}

func hasFoundIntergalactic(result func([]string) int) bool {
	return isValidCommand(result)
}

func isValidQues(command []string) (bool, func([]string) int) {
	var res func([]string) int
	if command[0] == "how" {
		howMuchIs := strings.Join(command[:3], " ")
		res = commandsMap[howMuchIs]
		if res == nil {
			howManyCreditsIs :=  strings.Join(command[:4], " ")
			res = commandsMap[howManyCreditsIs]
		}
		// return res != nil, res
	}else if command[0] == "Does" {
		res = commandsMap[command[0]]
	} else if command[0] == "Is" || strings.Contains(command[0], "Is") {
		res = commandsMap["Is"]
	}

	return res != nil, res
}
