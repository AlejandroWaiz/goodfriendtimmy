package cmdadapter

import (
	"fmt"
	"testing"

	timmyadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/TimmyAdapter"
)

func TestListenAndServe(t *testing.T) {

	testTable := []struct {
		caseName    string
		firstValue  int
		secondValue int
		operation   string
		expResult   int
		expError    error
	}{
		{"Normal add", 1, 2, "add", 3, nil},
		{"Normal subtract", 5, 3, "subtract", 2, nil},
		{"Normal multiply", 5, 3, "multiply", 15, nil},
		{"Normal divide", 10, 5, "divide", 2, nil},
		{"Divide by zero", 10, 0, "divide", 0, &timmyadapter.DivideByZero{}},
	}

	timmy := timmyadapter.CreateTimmyAdapter()

	cmdadap := CreateCmdAdapter(timmy)

	cmdadap.ListenAndServe()

	fmt.Println(cmdadap, testTable)

}
