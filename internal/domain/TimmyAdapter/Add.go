package timmyadapter

import (
	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
)

func (ta *TimmyAdapter) Add(operation domainstructs.Operation) domainstructs.Result {

	var result domainstructs.Result

	result.Is = operation.FirstOperand + operation.SecondOperand

	return result

}
