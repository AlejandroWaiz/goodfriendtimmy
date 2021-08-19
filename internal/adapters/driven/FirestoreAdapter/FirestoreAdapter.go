package firestoreadapter

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
)

type FirestoreAdapter struct {
	client *firestore.Client
}

var ctx context.Context

func CreateFirestoreAdapter() (*FirestoreAdapter, error) {

	ctx = context.Background()

	client, err := firestore.NewClient(ctx, os.Getenv("FirestoreProjectID"))

	if err != nil {
		return nil, err
	}

	return &FirestoreAdapter{client: client}, nil

}
