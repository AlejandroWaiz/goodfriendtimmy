package timmyadapter_test

import (
	"github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/TimmyAdapter"
	"testing"
)

func TestAdd(t *testing.T) {

	friend := timmyadapter.CreateTimmyAdapter()

	testTable := []struct {
		firstValue     int
		secondValue    int
		expectedResult int
		expectedError  error
	}{
		{firstValue: 5, secondValue: 5, expectedResult: 10, expectedError: nil},
	}

	for _, testCase := range testTable {

		obtainedResult := friend.Add(testCase.firstValue, testCase.secondValue)

		if obtainedResult != testCase.expectedResult {

			t.Errorf("Wrong result. Expected %v, got; %v ", testCase.expectedResult, obtainedResult)

		}

	}

}
