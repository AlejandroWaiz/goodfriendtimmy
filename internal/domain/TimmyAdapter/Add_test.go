package timmyadapter_test

import (
	"encoding/json"
	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
	"github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/TimmyAdapter"
	"log"
	"testing"
)

func TestAdd(t *testing.T) {

	friend := timmyadapter.CreateTimmyAdapter()

	testTable := []struct {
		domainstructs.Operation
		domainstructs.Result
		expectedError error
	}{
		{Operation: domainstructs.Operation{FirstOperand: 1, SecondOperand: 1}, Result: domainstructs.Result{Is: 2}, expectedError: nil},
		{Operation: domainstructs.Operation{FirstOperand: 2, SecondOperand: 2}, Result: domainstructs.Result{Is: 4}, expectedError: nil},
	}

	for _, testCase := range testTable {

		operation, err := json.Marshal(testCase.Operation)

		if err != nil {
			t.Log(err)
		}

		result, err := friend.Add(operation)

		if err != testCase.expectedError {
			t.Errorf("Expected %v error, got: %v", testCase.expectedError, err)
		}

		var finalResult domainstructs.Result
		err = json.Unmarshal(result, &finalResult)

		if err != nil {
			t.Log(err)
		}

		if finalResult.Is != testCase.Result.Is {
			t.Errorf("Expected %d result, got: %d", testCase.Result.Is, finalResult.Is)
		}

		log.Printf("Expected result was: %d, and we got: %d ", testCase.Result.Is, finalResult.Is)

	}

}
