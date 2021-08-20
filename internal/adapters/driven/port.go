package drivenadapters

import (
	firestorestructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/adapters/driven/FirestoreAdapter/structs"
)

type DrivenPort interface {
	Save(data firestorestructs.OperationToStore) (firestorestructs.OperationToStore, error)
	GetAll() ([]firestorestructs.OperationToStore, error)
}
