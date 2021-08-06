package domain

type DomainPort interface {
	Add([]byte) []byte
	Subtract([]byte) []byte
	Multiply([]byte) []byte
	Divide([]byte) ([]byte, error)
}
