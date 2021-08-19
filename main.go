package main

import (
	"os"

	cmdadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/adapters/driver/CmdAdapter"
	timmyadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/TimmyAdapter"
)

func main() {

	err := setEnvVariables()

	if err != nil {
		panic(err)
	}

	domainPort := timmyadapter.CreateTimmyAdapter()
	driverHandler := cmdadapter.CreateCmdAdapter(domainPort)
	driverHandler.ListenAndServe()

}

func setEnvVariables() error {

	err := os.Setenv("MuxPort", "3000")

	if err != nil {
		return err
	}

	err = os.Setenv("FirestoreProjectID", "ChangeThisValue")

	if err != nil {
		return err
	}

	return nil

}
