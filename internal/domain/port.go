package domain

type DomainPort interface {
	Add(firstValue, secondValue int) (result int)
	Subtract(firstValue, secondValue int) (result int)
	Multiply(firstValue, secondValue int) (result int)
	Divide(firstValue, secondValue int) (result int, err error)
}
