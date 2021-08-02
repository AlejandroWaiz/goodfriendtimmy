package timmyadapter_test

import (
	timmyadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/TimmyAdapter"
	"testing"
)

func TestSubtract(t *testing.T) {

	friend := timmyadapter.CreateTimmyAdapter()

	testTable := []struct {
		firstValue     int
		secondValue    int
		expectedResult int
	}{
		{firstValue: 5, secondValue: 5, expectedResult: 0},
	}

	for _, testCase := range testTable {

		obtainedResult := friend.Subtract(testCase.firstValue, testCase.secondValue)

		if obtainedResult != testCase.expectedResult {

			t.Errorf("Wrong result. Expected %v, got; %v", testCase.expectedResult, obtainedResult)

		}

	}

}
