package inputFile

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type FileOperations interface {
	open(filePath string) (*os.File, error)
	close(file *os.File) error
	read(file *os.File) [][]string
}

type fileOperations struct{}

func NewFileOperations() FileOperations {
	return &fileOperations{}
}

func (fo *fileOperations) open(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, errors.New("error opening the input file \n" +
			err.Error())
	}

	return file, nil
}

func (fo *fileOperations) close(file *os.File) error {
	return file.Close()
}

func (fo *fileOperations) read(file *os.File) [][]string{
	scanner := bufio.NewScanner(file)
	var argLists [][]string
	for scanner.Scan() {

		args := scanner.Text()
		argList := strings.Fields(args)
		argLists = append(argLists, argList)
	}

	return argLists
}