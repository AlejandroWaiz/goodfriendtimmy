package cmdadapter

import (
	"bufio"
	"os"
	"testing"

	timmyadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/TimmyAdapter"
)

func TestDoThisFunc(t *testing.T) {

	testTable := []struct {
		caseName      string
		operation     string
		FirstOperand  float64
		SecondOperand float64
		ExpResult     int
		ExpError      bool
	}{
		{"Normal add", "add", 1, 1, 2, false},
		{"Normal subtract", "subtract", 5, 3, 2, false},
		{"Normal Multiply", "multiply", 3, 5, 15, false},
		{"Normal divide", "divide", 6, 3, 2, false},
		{"Bad typing", "dvide", 6, 3, 0, true},
		{"Divide by zero", "divide", 10, 0, 0, true},
	}

	domain := timmyadapter.CreateTimmyAdapter()
	scanner := bufio.NewScanner(os.Stdin)

	cmd := CmdAdapter{port: domain, scanner: scanner}

	for _, testCase := range testTable {

		response, err := cmd.doThis(testCase.operation, testCase.FirstOperand, testCase.SecondOperand)

		if err != nil && testCase.ExpError == false {

			t.Errorf("Expected %v error, got %v on %v", testCase.ExpError, err, testCase.caseName)

		}

		if response != float64(testCase.ExpResult) {

			t.Errorf("Expected %v result, got %v on %v", testCase.ExpResult, response, testCase.caseName)

		}

	}

}
