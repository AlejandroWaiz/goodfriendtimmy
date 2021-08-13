package cmdadapter

import (
	"fmt"
)

func (cmd *CmdAdapter) ListenAndServe() {

	for {

		fmt.Println("Hello! Timmy here, i will help you with your math operation.")

		fmt.Println("What what would be the first value?")

		firstValue, err := cmd.comprobeInput(cmd.scanner.Text())

		if err != nil {
			fmt.Printf("Introduce a valid number")
			continue
		}

		fmt.Printf("%v right? nice, and the second one?", firstValue)

		secondValue, err := cmd.comprobeInput(cmd.scanner.Text())

		if err != nil {
			fmt.Printf("Introduce a valid number")
			continue
		}

		fmt.Println("Okey, and what would be the operation?")

		cmd.scanner.Scan()

		operation := cmd.scanner.Text()

		if operation != "add" && operation != "subtract" && operation != "multiply" && operation != "divide" {

			fmt.Printf("%v is not a valid operation.", operation)
			continue

		}

		result, err := cmd.doThis(operation, firstValue, secondValue)

		if err != nil {

			fmt.Printf("Sorry, i had this problem: %v", err)
			continue

		}

		fmt.Printf("Your result is %v !", result)

	}

}
