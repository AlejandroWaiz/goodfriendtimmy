package main

import (
	"log"
	"os"

	cmdadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/adapters/driver/CmdAdapter"
	timmyadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/TimmyAdapter"
)

func main() {

	err := os.Setenv("MuxPort", "3000")
	if err != nil {
		log.Fatalf("Could not set env variables: %v", err)
	}
	domainPort := timmyadapter.CreateTimmyAdapter()
	driverHandler := cmdadapter.CreateCmdAdapter(domainPort)
	driverHandler.ListenAndServe()

}
