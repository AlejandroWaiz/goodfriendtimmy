package cmdadapter

import (
	"fmt"
	"os"
)

func (cmd *CmdAdapter) ListenAndServe() {

	fmt.Println("what wea")

	for k := 0; k < 1; k++ {

		if i := 0; i == 0 {

			fmt.Println("Hello! Timmy here, i will help you with your math operation.")

			i++

		} else if i > 0 {

			fmt.Println("Lets do more operations!")

		}

		cmd.StartCalculatorLoop(os.Stdin)

	}
}
