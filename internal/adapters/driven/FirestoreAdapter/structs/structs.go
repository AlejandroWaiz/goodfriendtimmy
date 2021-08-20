package firestorestructs

import domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"

type OperationToStore struct {
	ProjectID   int64
	OperationIs string
	OperandsAre domainstructs.Operation
	Result      domainstructs.Result
}
