package timmyadapter

import "errors"

var divideByZero = errors.New("Cannot divide by 0, do you want to destroy the world?")

func (ta *TimmyAdapter) Divide(firstValue, secondValue int) (result int, err error) {

	if secondValue == 0 {

		return 0, divideByZero

	}

	result = firstValue / secondValue

	return result, nil

}
