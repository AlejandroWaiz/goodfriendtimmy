package HttpResponseMock

import (
	"encoding/json"
	"net/http"
)

type ResponseApi struct {
	Successs bool        `json:"Successs"`
	Status   int         `json:"status"`
	Result   interface{} `json:"result"`
}

func Success(result interface{}, status int) *ResponseApi {

	return &ResponseApi{
		Successs: true,
		Status:   status,
		Result:   result,
	}

}

func (r *ResponseApi) Send(w http.ResponseWriter) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	return json.NewEncoder(w).Encode(r)

}
