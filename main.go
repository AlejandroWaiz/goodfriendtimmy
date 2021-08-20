package main

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	firestoreadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/adapters/driven/FirestoreAdapter"
	muxadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/adapters/driver/HttpAdapter"
	timmyadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/TimmyAdapter"
	"github.com/joho/godotenv"
)

func main() {

	client, err := setEnvVariables()

	if err != nil {
		panic(err)
	}

	domainAdapter := timmyadapter.CreateTimmyAdapter()
	drivenHandler, err := firestoreadapter.CreateFirestoreAdapter(client)

	if err != nil {
		panic(err)
	}

	driverHandler := muxadapter.CreateMuxAdapter(domainAdapter, drivenHandler)
	driverHandler.ListenAndServe()

}

func setEnvVariables() (client *firestore.Client, err error) {

	err = godotenv.Load()

	c, err := firestore.NewClient(context.Background(), os.Getenv("goodfriendtimmy"))

	if err != nil {
		return
	}

	return c, nil

}
