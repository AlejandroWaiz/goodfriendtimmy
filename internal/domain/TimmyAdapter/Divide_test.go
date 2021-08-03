package timmyadapter_test

import (
	timmyadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/TimmyAdapter"
	"testing"
)

func TestDivide(t *testing.T) {

	friend := timmyadapter.CreateTimmyAdapter()

	testTable := []struct {
		firstValue     int
		secondValue    int
		expectedResult int
		expectedError  error
	}{
		{firstValue: 5, secondValue: 5, expectedResult: 1, expectedError: nil},
		{firstValue: 5, secondValue: 0, expectedResult: 0, expectedError: timmyadapter.DivideByZero},
	}

	for _, testCase := range testTable {

		obtainedResult, err := friend.Divide(testCase.firstValue, testCase.secondValue)

		if err != testCase.expectedError {

			t.Errorf("Expected %v error, got; %v", testCase.expectedError, err)

		}

		if obtainedResult != testCase.expectedResult {

			t.Errorf("Wrong result. Expected %v , got; %v", testCase.expectedResult, obtainedResult)

		}

	}
}
