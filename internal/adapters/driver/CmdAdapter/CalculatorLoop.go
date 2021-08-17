package cmdadapter

import (
	"fmt"
	"io"
	"log"

	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
)

func (cmd *CmdAdapter) StartCalculatorLoop(readOp, firstValue, secondValue io.Reader) (float64, error) {

	var operation domainstructs.Operation

	fmt.Println("What would be the first value?")

	var err error

	operation.FirstOperand, err = comprobeNumberInput(firstValue)

	if err != nil {
		fmt.Printf("Introduce a valid number")
		return 0, err
	}

	fmt.Printf("%v right? nice, and the second one?\n", operation.FirstOperand)

	operation.SecondOperand, err = comprobeNumberInput(secondValue)

	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	fmt.Println("Okey, and what would be the operation?")

	mathOperation := comprobeStringInput(readOp)

	fmt.Printf("%v, right? Lets see...\n", mathOperation)

	var finalresult domainstructs.Result

	finalresult.Is, err = cmd.doThis(mathOperation, operation.FirstOperand, operation.SecondOperand)

	return finalresult.Is, err

}
