package muxadapter

import (
	"github.com/AlejandroWaiz/goodfriendtimmy/infrastructure/HttpResponseMock"
	"io/ioutil"
	"net/http"
)

func (ma *MuxAdapter) Add(w http.ResponseWriter, r *http.Request) []byte {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {

		HttpResponseMock.ErrInvalidJSON.Send(w)

	}

	return body

}
