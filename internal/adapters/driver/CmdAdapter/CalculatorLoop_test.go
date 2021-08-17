package cmdadapter

import (
	"fmt"
	"strings"
	"testing"

	timmyadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/TimmyAdapter"
)

func TestStartCalculatorLoop(t *testing.T) {

	domain := timmyadapter.CreateTimmyAdapter()
	cmd := CreateCmdAdapter(domain)

	type operation struct {
		firstNumber  string
		secondNumber string
		operation    string
	}

	testTable := []struct {
		op        operation
		name      string
		expErr    bool
		expResult float64
	}{
		{name: "Normal add", expResult: 11, op: operation{firstNumber: "5", secondNumber: "6", operation: "add"}},
		{name: "Normal multiply", expResult: 15, op: operation{firstNumber: "5", secondNumber: "3", operation: "multiply"}},
		{name: "Normal subtract", expResult: 0, op: operation{firstNumber: "5", secondNumber: "5", operation: "subtract"}},
		{name: "Normal divide", expResult: 2, op: operation{firstNumber: "10", secondNumber: "5", operation: "divide"}},
		{name: "Bad operation typing", expResult: 0, expErr: true, op: operation{firstNumber: "10", secondNumber: "5", operation: "dvide"}},
		{name: "Bad number typing", expResult: 0, expErr: true, op: operation{firstNumber: "10sd", secondNumber: "5", operation: "divide"}},
		{name: "Divide by zero", expResult: 0, expErr: true, op: operation{firstNumber: "10", secondNumber: "0", operation: "divide"}},
	}
	for _, testCase := range testTable {

		t.Run(testCase.name, func(t *testing.T) {

			operation := strings.NewReader(testCase.op.operation)
			firstValue := strings.NewReader(testCase.op.firstNumber)
			secondValue := strings.NewReader(testCase.op.secondNumber)

			fmt.Printf("operation: %T, %v. first: %T, %v.second: %T, %v.", operation, operation, firstValue, firstValue, secondValue, secondValue)

			result, err := cmd.StartCalculatorLoop(operation, firstValue, secondValue)

			if result != testCase.expResult {
				t.Errorf("Want %v as result, got; %v in %v", testCase.expResult, result, testCase.name)
			}

			if err != nil && testCase.expErr == false {
				t.Errorf("Doesnt expecting error got; %v in %v", err, testCase.name)
			}

		})
	}

}
