package inputFile

import (
	"galaxy_merchant_trading_guide/src/commandLine"
	"os"
	"strings"
	"testing"
)

func TestOpenNonExistentFile(t *testing.T) {
	fo := NewFileOperations()
	nonExistentFilePath := "nonexistentfile.txt"
	_, openErr := fo.open(nonExistentFilePath)
	if openErr == nil {
		t.Error("Expected an error, but got nil")
	}
	expectedErrMsg := "error opening the input file"
	if !strings.Contains(openErr.Error(), expectedErrMsg) {
		t.Errorf("Expected error: %v, but got: %v", expectedErrMsg, openErr)
	}
}

func TestOpen(t *testing.T) {
	file, err := os.CreateTemp("", "testfile*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	fo := NewFileOperations()
	filePath := file.Name()
	openedFile, openErr := fo.open(filePath)
	if openErr != nil {
		t.Errorf("Expected no error, but got: %v", openErr)
	}
	defer openedFile.Close() 

}

func TestClose(t *testing.T) {
	file, err := os.CreateTemp("", "testfile*.txt")
	if err != nil {
		t.Fatal(err)
	}

	fo := NewFileOperations()

	fo.close(file)

	_, readErr := file.Read([]byte{})
	if readErr == nil {
		t.Error("Expected the file to be closed, but it's not")
	}
}

func TestRead(t *testing.T) {
	file, err := os.CreateTemp("", "testfile*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	lines := []string{
		"command1 arg1 arg2",
		"command2 arg3",
		"command3",
	}
	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			t.Fatal(err)
		}
	}
	file.Close()

	fo := NewFileOperations()
	file, err = fo.open(file.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer fo.close(file)
	

	argLists := fo.read(file)

	expectedLines := len(lines)
	actualLines := len(argLists)
	if actualLines != expectedLines {
		t.Errorf("Expected %d lines, but got %d", expectedLines, actualLines)
	}

	for i, line := range lines {
		expectedArgs := strings.Fields(line)
		actualArgs := argLists[i]

		if !commandLine.EqualStringSlices(expectedArgs, actualArgs) {
			t.Errorf("Line %d - Expected: %v, but got: %v", i+1, expectedArgs, actualArgs)
		}
	}

}
