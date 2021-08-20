package firestoreadapter

import (
	firestorestructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/adapters/driven/FirestoreAdapter/structs"
	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
	"google.golang.org/api/iterator"
)

func (fa *FirestoreAdapter) GetAll() ([]firestorestructs.OperationToStore, error) {

	var operations []firestorestructs.OperationToStore

	iter := fa.client.Collection("Operations").Documents(ctx)

	for {

		doc, err := iter.Next()

		if err == iterator.Done {

			break

		}

		if err != nil {

			return nil, err

		}

		operation := firestorestructs.OperationToStore{
			ProjectID:   doc.Data()["ProyectID"].(int64),
			OperationIs: doc.Data()["OperationType"].(string),
			OperandsAre: domainstructs.Operation{
				FirstOperand:  doc.Data()["FirstOperand"].(int64),
				SecondOperand: doc.Data()["SecondOperand"].(int64)},
			Result: domainstructs.Result{Is: doc.Data()["OperationResult"].(int64)},
		}

		operations = append(operations, operation)

	}

	return operations, nil

}
