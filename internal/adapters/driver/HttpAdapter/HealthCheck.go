package muxadapter

import (
	"fmt"
	"log"
	"net/http"
)

func (ma *MuxAdapter) HealthCheck(w http.ResponseWriter, r *http.Request) {

	_, err := fmt.Fprint(w, "Hello there!")

	if err != nil {
		w.WriteHeader(400)
		log.Fatal(err)
	}

}
