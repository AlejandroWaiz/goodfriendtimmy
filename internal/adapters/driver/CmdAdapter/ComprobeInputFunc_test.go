package cmdadapter

import (
	"strings"
	"testing"
)

func TestComprobeNumberInput(t *testing.T) {

	testCases := []struct {
		text     string
		expValue int64
		expErr   bool
	}{
		{text: "5", expErr: false, expValue: 5},
		{text: "Five", expErr: true},
	}

	for _, testthis := range testCases {

		readThis := strings.NewReader(testthis.text)

		obtainedValue, err := comprobeNumberInput(readThis)

		if err != nil && testthis.expErr == false {
			t.Log(err)
		}

		if obtainedValue != testthis.expValue {
			t.Error(err)
		}

	}

}

func TestComprobeStringInput(t *testing.T) {

	testCases := []struct {
		operation string
		expResult string
		expErr    bool
	}{
		{operation: "add", expErr: false, expResult: "add"},
		{operation: "Five", expErr: false, expResult: "Five"},
		{operation: "", expErr: true},
	}

	for _, testthis := range testCases {

		readThis := strings.NewReader(testthis.operation)

		obtainedValue := comprobeStringInput(readThis)

		if obtainedValue != testthis.expResult {
			t.Errorf("Bad result, expected :%v, got: %v", testthis.expResult, obtainedValue)
		}

	}

}
