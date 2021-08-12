package timmyadapter_test

import (
	"testing"

	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
	timmyadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/TimmyAdapter"
)

func TestAdd(t *testing.T) {

	friend := timmyadapter.CreateTimmyAdapter()

	testTable := []struct {
		domainstructs.Operation
		ExpectedResult domainstructs.Result
		expectedError  error
	}{
		{Operation: domainstructs.Operation{FirstOperand: 1, SecondOperand: 1}, ExpectedResult: domainstructs.Result{Is: 2}, expectedError: nil},
		{Operation: domainstructs.Operation{FirstOperand: 2, SecondOperand: 2}, ExpectedResult: domainstructs.Result{Is: 4}, expectedError: nil},
	}

	for _, testCase := range testTable {

		result := friend.Add(testCase.Operation)

		if result != testCase.ExpectedResult {
			t.Fatalf("Expected %d as a result, got: %d", testCase.ExpectedResult, result)
		}

	}

}
