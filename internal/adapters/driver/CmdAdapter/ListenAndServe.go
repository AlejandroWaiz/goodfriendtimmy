package cmdadapter

import (
	"fmt"
)

func (cmd *CmdAdapter) ListenAndServe() {

	for {

		if i := 0; i == 0 {

			fmt.Println("Hello! Timmy here, i will help you with your math operation.")

			i++

		} else if i > 0 {

			fmt.Println("Lets do more operations!")

		}

		fmt.Println("What what would be the first value?")

		firstValue, err := cmd.comprobeInput()

		if err != nil {
			fmt.Printf("Introduce a valid number")
			continue
		}

		fmt.Printf("%v right? nice, and the second one?\n", firstValue)

		secondValue, err := cmd.comprobeInput()

		if err != nil {
			fmt.Println("Introduce a valid number")
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

			fmt.Printf("Sorry, i had this problem: %v\n", err)
			continue

		}

		fmt.Printf("Your result is %v !\n", result)

	}

}
