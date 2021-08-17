package cmdadapter

import (
	"fmt"
	"io"
	"log"

	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
)

func (cmd *CmdAdapter) StartCalculatorLoop(fromHere io.Reader) {

	var operation domainstructs.Operation

	fmt.Println("What what would be the first value?")

	var err error

	operation.FirstOperand, err = comprobeNumberInput(fromHere)

	if err != nil {
		fmt.Printf("Introduce a valid number")
		return
	}

	fmt.Printf("%v right? nice, and the second one?\n", operation.FirstOperand)

	operation.SecondOperand, err = comprobeNumberInput(fromHere)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Okey, and what would be the operation?")

	mathOperation, err := comprobeStringInput(fromHere)

	if err != nil {
		log.Fatal(err)
		return
	}

	var finalresult domainstructs.Result

	finalresult.Is, err = cmd.doThis(mathOperation, operation.FirstOperand, operation.SecondOperand)

	if err != nil {

		fmt.Printf("Sorry, i had this problem: %v\n", err)
		return
	}

	fmt.Printf("Your result is %v !\n", finalresult.Is)
	fmt.Printf("%+v", finalresult)

}
