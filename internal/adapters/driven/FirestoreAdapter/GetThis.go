package firestoreadapter

import (
	"os"

	firestorestructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/adapters/driven/FirestoreAdapter/structs"
	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
	"google.golang.org/api/iterator"
)

func (fa *FirestoreAdapter) GetAll() ([]firestorestructs.OperationToStore, error) {

	defer fa.client.Close()

	var operations []firestorestructs.OperationToStore
	iter := fa.client.Collection(os.Getenv("Firestore-Collection")).Documents(ctx)

	defer iter.Stop()

	for {

		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}

		operation := firestorestructs.OperationToStore{
			ProjectID:   doc.Data()["Proyect ID"].(int),
			OperationIs: doc.Data()["Operation Type"].(string),
			OperandsAre: domainstructs.Operation{
				FirstOperand:  doc.Data()["First Operand"].(float64),
				SecondOperand: doc.Data()["Second Operand"].(float64)},
			ResultIs: domainstructs.Result{Is: doc.Data()["Operation Result"].(float64)},
		}

		operations = append(operations, operation)

	}

	return operations, nil

}
