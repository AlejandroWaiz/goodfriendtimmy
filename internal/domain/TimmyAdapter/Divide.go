package timmyadapter

import (
	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
)

func (ta *TimmyAdapter) Divide(operation domainstructs.Operation) (result domainstructs.Result, err error) {

	if operation.SecondOperand == 0 {

		return result, &DivideByZero{}

	}

	result.Is = int64(operation.FirstOperand) / int64(operation.SecondOperand)

	return result, nil

}
