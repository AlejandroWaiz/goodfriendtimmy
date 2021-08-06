package timmyadapter_test

import (
	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
	timmyadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/TimmyAdapter"
	"testing"
)

func TestDivide(t *testing.T) {

	friend := timmyadapter.CreateTimmyAdapter()

	testTable := []struct {
		Operation      domainstructs.Operation
		ExpectedResult domainstructs.Result
		expectedError  error
	}{
		{Operation: domainstructs.Operation{FirstOperand: 5, SecondOperand: 5}, ExpectedResult: domainstructs.Result{Is: 1}, expectedError: nil},
		{Operation: domainstructs.Operation{FirstOperand: 5, SecondOperand: 0}, ExpectedResult: domainstructs.Result{Is: 0}, expectedError: &timmyadapter.DivideByZero{}},
	}

	for _, testCase := range testTable {

		result, err := friend.Divide(testCase.Operation)

		if err != testCase.expectedError {
			t.Fatalf("Expected %v error, got: %v", testCase.expectedError, err)
		}

		if result != testCase.ExpectedResult {
			t.Fatalf("Expected %d as a result, got: %d", testCase.ExpectedResult, result)
		}

	}
}
