package muxadapter

import (
	"encoding/json"
	web "github.com/AlejandroWaiz/goodfriendtimmy/infrastructure/HttpResponseMock"
	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
	"io/ioutil"
	"log"
	"net/http"
)

func (ma *MuxAdapter) DivideHttpHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {

		web.ErrInvalidJSON.Send(w)

	}

	var operation domainstructs.Operation

	err = json.Unmarshal(body, &operation)

	if err != nil {
		log.Printf("Sorry, got this err: %v", err)
	}

	result, err := ma.domainport.Divide(operation)

	if err != nil {
		web.ErrBadRequest.Send(w)
		log.Println(err)
	}

	web.Success(result, 200)

}
