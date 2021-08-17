package cmdadapter

import (
	"fmt"
	"os"
)

func (cmd *CmdAdapter) ListenAndServe() {

	for k := 0; k < 1; k++ {

		if i := 0; i == 0 {

			fmt.Println("Hello! Timmy here, i will help you with your math operation.")

			i++

		} else if i > 0 {

			fmt.Println("Lets do more operations!")

		}

		result, err := cmd.StartCalculatorLoop(os.Stdin, os.Stdin, os.Stdin)

		if err != nil {
			fmt.Printf("Sorry, i had this problem: %v\n", err)
			return
		}

		fmt.Printf("Your result is %v!\n", result)

	}
}
