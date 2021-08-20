package muxadapter

import (
	"encoding/json"
	"net/http"

	"github.com/AlejandroWaiz/goodfriendtimmy/infrastructure/HttpResponseMock"
)

func (ma *MuxAdapter) GetAllSavedOperations(w http.ResponseWriter, r *http.Request) {

	data, err := ma.drivenport.GetAll()

	if err != nil {
		HttpResponseMock.InternalError.Send(w)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)

}
