package muxadapter

import (
	"encoding/json"
	web "github.com/AlejandroWaiz/goodfriendtimmy/infrastructure/HttpResponseMock"
	"log"

	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
	"io/ioutil"
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
		web.ErrInvalidJSON.Send(w)
	}

	result := ma.domainport.Add(operation)

	err = json.NewEncoder(w).Encode(result)
	log.Printf("The result in the real func is: %v", result.Is)

	if err != nil {
		web.InternalError.Send(w)
	}

	//

}
