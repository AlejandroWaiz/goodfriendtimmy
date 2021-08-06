package timmyadapter

type DivideByZero struct{}

func (e *DivideByZero) Error() string {
	return "Cant divide by zero... Do you want to destroy the universe!?."
}

type WrongBodyForOperation struct{}

func (e *WrongBodyForOperation) Error() string {
	return "Wrong body for operation."
}
