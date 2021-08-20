package main

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	cmdadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/adapters/driver/CmdAdapter"
	timmyadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/TimmyAdapter"
	"google.golang.org/api/option"
)

func main() {

	_, err := setEnvVariables()

	if err != nil {
		panic(err)
	}

	domainPort := timmyadapter.CreateTimmyAdapter()
	driverHandler := cmdadapter.CreateCmdAdapter(domainPort)
	driverHandler.ListenAndServe()

}

func setEnvVariables() (client *firestore.Client, err error) {

	err = os.Setenv("MuxPort", "3000")

	if err != nil {
		return
	}

	err = os.Setenv("FirestoreProjectID", "ChangeThisValue")

	if err != nil {
		return
	}

	//Variables here are used if and only if firestore will be used. If dont, dont use ''Client'' return vale in main func.
	sa := option.WithCredentialsFile("./ServiceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)

	if err != nil {
		return
	}

	client, err = app.Firestore(context.Background())

	if err != nil {
		return
	}

	return client, nil

}
