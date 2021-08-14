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
		FirstOperand  int
		SecondOperand int
		ExpResult     int
		ExpError      error
	}{
		{"Normal add", "add", 1, 1, 2, nil},
		{"Normal subtract", "subtract", 5, 3, 2, nil},
		{"Normal Multiply", "multiply", 3, 5, 15, nil},
		{"Normal divide", "divide", 6, 3, 2, nil},
		{"Divide by zero", "divide", 10, 0, 10, &timmyadapter.DivideByZero{}},
	}

	domain := timmyadapter.CreateTimmyAdapter()
	scanner := bufio.NewScanner(os.Stdin)

	cmd := CmdAdapter{port: domain, scanner: scanner}

	for _, testCase := range testTable {

		response, err := cmd.doThis(testCase.operation, testCase.FirstOperand, testCase.SecondOperand)

		if err != testCase.ExpError {

			t.Errorf("Expected %v error, got %v on %v", testCase.ExpError, err, testCase.caseName)

		}

		if response != float64(testCase.ExpResult) {

			t.Errorf("Expected %v result, got %v on %v", testCase.ExpResult, response, testCase.caseName)

		}

	}

}
