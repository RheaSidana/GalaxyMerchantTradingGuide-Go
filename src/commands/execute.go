package commands

import (
	"fmt"
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
	result := commandsMap[command[0]]

	if !isValidCommand(result) {
		// return errors.New(INPUT_DATA_ERROR)
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