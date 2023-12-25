package commandLine

import (
	"errors"
)

func ParseCommandLineArgs() (string, error) {
	cliArgs, err := Arguments()
	if err != nil {
		return "", err
	}
	return extractFilePath(cliArgs)
}

func extractFilePath(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("missing file path argument")
	}
	return args[0], nil
}