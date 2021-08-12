package muxadapter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	web "github.com/AlejandroWaiz/goodfriendtimmy/infrastructure/HttpResponseMock"
	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
)

func (ma *MuxAdapter) SubtractHttpHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {

		web.ErrInvalidJSON.Send(w)

	}

	var operation domainstructs.Operation

	err = json.Unmarshal(body, &operation)

	if err != nil {
		log.Printf("Sorry, got this err: %v", err)
	}

	result := ma.domainport.Subtract(operation)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(result)

}
