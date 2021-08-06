package timmyadapter

import (
	"encoding/json"
	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
)

func (ta *TimmyAdapter) Subtract(data []byte) ([]byte, error) {

	var operation domainstructs.Operation

	err := json.Unmarshal(data, &operation)

	if err != nil {

		return nil, &WrongBodyForOperation{}

	}

	var result domainstructs.Result

	result.Is = operation.FirstOperand - operation.SecondOperand

	resp, err := json.Marshal(result)

	if err != nil {

		return nil, err

	}

	return resp, nil

}
