package firestorestructs

import domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"

type OperationToStore struct {
	ProjectID   int
	OperationIs string
	OperandsAre domainstructs.Operation
	ResultIs    domainstructs.Result
}
