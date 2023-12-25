package inputFile

import (
	"errors"
	"os"
	"testing"
)

type mockFileOperations struct {
	openFunc  func(filePath string) (*os.File, error)
	closeFunc func(file *os.File) error
	readFunc  func(file *os.File) [][]string
}

func (m *mockFileOperations) open(filePath string) (*os.File, error) {
	if m.openFunc != nil {
		return m.openFunc(filePath)
	}
	return nil, errors.New("mocked open error")
}

func (m *mockFileOperations) close(file *os.File) error {
	return m.closeFunc(file)
}

func (m *mockFileOperations) read(file *os.File) [][]string {
	if m.readFunc != nil {
		return m.readFunc(file)
	}
	return nil
}

func TestReadCommandsWhenOpenReturnsError(t *testing.T) {
	mockFO := new(mockFileOperations)
	mockFO.openFunc = func(filePath string) (*os.File, error) {
		return nil, errors.New("mocked open error")
	}
	// mockFO.closeFunc = func(file *os.File) error{
	// 	return nil
	// }
	// mockFO.readFunc = func(file *os.File) [][]string {
	// 	return [][]string{{"command1", "arg1"}, {"command2", "arg2"}}
	// }
	fileCommands := fileCommands{fileOperations: mockFO}

	filePath := "testfile.txt"
	argLists, err := fileCommands.ReadCommands(filePath)

	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if argLists != nil {
		t.Errorf("Expected argLists to be nil, but got: %v", argLists)
	}
}

func TestReadCommands(t *testing.T) {
	mockFO := new(mockFileOperations)
	mockFO.openFunc = func(filePath string) (*os.File, error) {
		return &os.File{}, nil
	}
	mockFO.closeFunc = func(file *os.File) error {
		return errors.New("mocked open error")
	}
	mockFO.readFunc = func(file *os.File) [][]string {
		return [][]string{{"command1", "arg1"}, {"command2", "arg2"}}
	}
	fileCommands := fileCommands{fileOperations: mockFO}

	filePath := "testfile.txt"
	argLists, err := fileCommands.ReadCommands(filePath)

	if err != nil {
		t.Error("Expected not an error")
	}
	if argLists == nil {
		t.Error("Expected argLists to be not nil")
	}

}
