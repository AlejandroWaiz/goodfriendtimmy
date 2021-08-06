package muxadapter

import (
	"fmt"
	"net/http"
)

func (ma *MuxAdapter) HealthCheck(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello there!")

}
