package muxadapter

import (
	"encoding/json"
	web "github.com/AlejandroWaiz/goodfriendtimmy/infrastructure/HttpResponseMock"
	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
	"io/ioutil"
	"log"
	"net/http"
)

func (ma *MuxAdapter) AddHttpHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {

		web.ErrInvalidJSON.Send(w)

	}

	var operation domainstructs.Operation

	err = json.Unmarshal(body, &operation)

	if err != nil {
		log.Printf("Sorry, got this err: %v", err)
	}

	result := ma.domainport.Add(operation)

	web.Success(result, 200)

}
