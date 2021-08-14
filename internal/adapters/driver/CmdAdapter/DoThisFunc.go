package cmdadapter

import domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"

func (cmd *CmdAdapter) doThis(operation string, firstValue int, secondValue int) (response float64, err error) {

	operands := domainstructs.Operation{FirstOperand: float64(firstValue), SecondOperand: float64(secondValue)}

	if operation == "add" {

		result := cmd.port.Add(operands)
		return result.Is, nil

	} else if operation == "subtract" {

		result := cmd.port.Subtract(operands)
		return result.Is, nil

	} else if operation == "multiply" {

		result := cmd.port.Multiply(operands)
		return result.Is, nil

	} else if operation == "divide" {

		result, err := cmd.port.Divide(operands)

		if err != nil {
			return result.Is, err
		}

		return result.Is, nil

	}

	return

}
