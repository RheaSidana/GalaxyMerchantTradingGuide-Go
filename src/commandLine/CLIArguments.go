package commandLine

import (
	"errors"
	"os"
)

func Arguments() ([]string, error) {
	fileNumStarts := 1
	cliArgs := os.Args[fileNumStarts:]

	if len(cliArgs) == 0 {
		return nil, errors.New("please provide the input file path")
	}

	return cliArgs, nil
}

func EqualStringSlices(slice1 []string, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}