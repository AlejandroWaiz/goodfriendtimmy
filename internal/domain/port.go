package domain

import domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"

type DomainPort interface {
	Add(domainstructs.Operation) domainstructs.Result
	Subtract(domainstructs.Operation) domainstructs.Result
	Multiply(domainstructs.Operation) domainstructs.Result
	Divide(domainstructs.Operation) (domainstructs.Result, error)
}
