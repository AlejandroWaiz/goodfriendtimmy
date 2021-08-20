package firestoreadapter

import (
	"context"

	"cloud.google.com/go/firestore"
)

type FirestoreAdapter struct {
	client *firestore.Client
}

var ctx context.Context

func CreateFirestoreAdapter(client *firestore.Client) (*FirestoreAdapter, error) {

	ctx = context.Background()

	return &FirestoreAdapter{client: client}, nil

}
