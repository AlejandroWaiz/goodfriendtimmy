package timmyadapter

import (
	"log"

	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
)

func (ta *TimmyAdapter) Divide(operation domainstructs.Operation) (result domainstructs.Result, err error) {

	if operation.SecondOperand == 0 {

		return result, &DivideByZero{}

	}

	result.Is = float64(operation.FirstOperand) / float64(operation.SecondOperand)
	log.Printf("The result is %v", result.Is)

	return result, nil

}
