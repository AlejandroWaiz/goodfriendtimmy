package timmyadapter

import (
	"log"

	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
)

func (ta *TimmyAdapter) Subtract(operation domainstructs.Operation) (result domainstructs.Result) {

	result.Is = operation.FirstOperand - operation.SecondOperand
	log.Printf("The result is: %v", result)

	return result

}
