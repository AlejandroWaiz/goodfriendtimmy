package main

import (
	muxadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/adapters/driver/HttpAdapter"
	timmyadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/TimmyAdapter"
	"log"
	"os"
)

func main() {

	err := os.Setenv("MuxPort", "3000")

	if err != nil {

		log.Fatalf("Could not set env variables: %v", err)

	}

	domainPort := timmyadapter.CreateTimmyAdapter()

	MuxPort := muxadapter.CreateMuxAdapter(domainPort)

	MuxPort.ListenAndServe()

}
