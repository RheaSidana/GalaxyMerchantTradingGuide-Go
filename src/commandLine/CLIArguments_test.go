package commandLine

import (
	"os"
	"testing"
)

func TestArgumentsWhenNoArgsAreProvided(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{"test-program"}
	// the below statement will produce error!=nil and result==nil
	result, err := Arguments()
	if err == nil {
		t.Error("Expected error: nil")
	}
	if result != nil {
		t.Errorf("Expected result to be nil, but got: %v", result)
	}
}

func TestArguments(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{"test-program", "input.txt", "--flag", "value"}
	result, err := Arguments()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	expectedResult := []string{"input.txt", "--flag", "value"}
	if !EqualStringSlices(result, expectedResult) {
		t.Errorf("Expected result: %v, but got: %v", expectedResult, result)
	}
}
