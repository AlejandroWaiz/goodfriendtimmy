package main

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	firestoreadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/adapters/driven/FirestoreAdapter"
	muxadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/adapters/driver/HttpAdapter"
	timmyadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/TimmyAdapter"
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

	err = os.Setenv("MuxPort", "3000")

	if err != nil {
		return
	}

	err = os.Setenv("FirestoreProjectID", "goodfriendtimmy")

	err = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/Users/alejandrowaiz/Desktop/goodfriendtimmy/ServiceAccountKey.json")

	if err != nil {
		return
	}

	//Variables here are used if and only if firestore will be used. If dont, dont use ''Client'' return vale in main func.
	//sa := option.WithCredentialsFile("/Users/alejandrowaiz/Desktop/goodfriendtimmy/ServiceAccountKey.json")
	c, _ := firestore.NewClient(context.Background(), "goodfriendtimmy")

	if err != nil {
		return
	}

	return c, nil

	//	client, err = app.Firestore(context.Background())

	//	if err != nil {
	//		return
	//	}

	//	return client, nil

}
