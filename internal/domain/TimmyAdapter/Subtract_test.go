package timmyadapter_test

import (
	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
	timmyadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/TimmyAdapter"
	"testing"
)

func TestSubtract(t *testing.T) {

	friend := timmyadapter.CreateTimmyAdapter()

	testTable := []struct {
		Operation      domainstructs.Operation
		ExpectedResult domainstructs.Result
	}{
		{Operation: domainstructs.Operation{FirstOperand: 5, SecondOperand: 5}, ExpectedResult: domainstructs.Result{Is: 0}},
		{Operation: domainstructs.Operation{FirstOperand: 5, SecondOperand: 0}, ExpectedResult: domainstructs.Result{Is: 5}},
	}

	for _, testCase := range testTable {

		result := friend.Subtract(testCase.Operation)

		if result != testCase.ExpectedResult {
			t.Fatalf("Expected %d as a result, got: %d", testCase.ExpectedResult, result)
		}

	}
}
