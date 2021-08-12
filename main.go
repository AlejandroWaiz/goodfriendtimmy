package main

import (
	"log"

	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
	"github.com/goccy/go-json"
)

func main() {

	/* err := os.Setenv("MuxPort", "3000")

	if err != nil {

		log.Fatalf("Could not set env variables: %v", err)

	}

	domainPort := timmyadapter.CreateTimmyAdapter()

	MuxPort := muxadapter.CreateMuxAdapter(domainPort)

	MuxPort.ListenAndServe() */

	var result domainstructs.Result = domainstructs.Result{Is: 3}

	finalresult, err := json.Marshal(result)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(finalresult)
}
