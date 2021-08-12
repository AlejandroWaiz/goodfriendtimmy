package muxadapter

import (
	"encoding/json"

	web "github.com/AlejandroWaiz/goodfriendtimmy/infrastructure/HttpResponseMock"

	"io/ioutil"
	"net/http"

	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
)

func (ma *MuxAdapter) AddHttpHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {

		web.ErrInvalidJSON.Send(w)
		return

	}

	var operation domainstructs.Operation

	err = json.Unmarshal(body, &operation)

	if err != nil {

		web.ErrInvalidJSON.Send(w)
		return

	}

	result := ma.domainport.Add(operation)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(result)

	/*err = json.NewEncoder(w).Encode(result)
	if err != nil {

		xerr := web.InternalError.Send(w)

		if xerr != nil {
			w.WriteHeader(500)
		}
	} */

}
