package firestoreadapter

import (
	"os"

	firestorestructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/adapters/driven/FirestoreAdapter/structs"
)

func (fa *FirestoreAdapter) Save(data firestorestructs.OperationToStore) (firestorestructs.OperationToStore, error) {

	defer fa.client.Close()
	_, _, err := fa.client.Collection(os.Getenv("Firestore-Collection")).Add(ctx, map[string]interface{}{
		"ID":               data.ProjectID,
		"Operation Type":   data.OperationIs,
		"First Operand":    data.OperandsAre.FirstOperand,
		"Second Operand":   data.OperandsAre.SecondOperand,
		"Operation Result": data.ResultIs,
	})

	if err != nil {
		return data, err
	}

	return data, nil

}
