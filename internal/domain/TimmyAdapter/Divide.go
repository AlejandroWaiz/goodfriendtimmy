package timmyadapter

import (
	"encoding/json"
	"errors"
	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
)

var divideByZero = errors.New("Cannot divide by 0, do you want to destroy the world?")

func (ta *TimmyAdapter) Divide(data []byte) ([]byte, error) {

	var operation domainstructs.Operation

	err := json.Unmarshal(data, &operation)

	if err != nil {
		return nil, &WrongBodyForOperation{}
	}

	if operation.SecondOperand == 0 {
		return nil, &DivideByZero{}
	}

	var result domainstructs.Result

	result.Is = operation.FirstOperand / operation.SecondOperand

	resp, err := json.Marshal(result)

	if err != nil {
		return nil, err
	}

	return resp, nil

}
