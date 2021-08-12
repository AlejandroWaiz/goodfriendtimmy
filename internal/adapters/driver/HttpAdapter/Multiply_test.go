package muxadapter

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
	timmyadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/TimmyAdapter"
)

func TestMultiply(t *testing.T) {

	testTable := []struct {
		FirstOperand  float64 `json:",omitempty"`
		SecondOperand float64 `json:",omitempty"`
		ExpResult     float64
		ExpHttpStatus int
		ExpHttpHeader string
		ExpBody       domainstructs.Result
	}{
		{FirstOperand: 1, SecondOperand: 2, ExpResult: 2, ExpHttpStatus: 200, ExpHttpHeader: "application/json", ExpBody: domainstructs.Result{Is: 2}},
		{FirstOperand: 1, ExpHttpHeader: "application/json", ExpBody: domainstructs.Result{Is: 0}, ExpResult: 0, ExpHttpStatus: 200},
	}

	for i, testCase := range testTable {

		opbody := domainstructs.Operation{FirstOperand: testCase.FirstOperand, SecondOperand: testCase.SecondOperand}

		bodyInBytes, err := json.Marshal(opbody)

		if err != nil {
			t.Log(err)
		}

		buf := bytes.NewBuffer(bodyInBytes)

		if err != nil {
			t.Log(err)
		}

		request, err := http.NewRequest("POST", "http://localhost:8080/Multiply", buf)
		request.Header.Set("Content-type", "application/json")

		if err != nil {
			t.Log(err)
		}

		recorder := httptest.NewRecorder()

		timmy := timmyadapter.CreateTimmyAdapter()
		muxAdapter := MuxAdapter{domainport: timmy}

		muxAdapter.MultiplyHttpHandler(recorder, request)

		response := recorder.Result()

		httpbody, err := ioutil.ReadAll(response.Body)

		if err != nil {

			t.Fatalf("Could not read response: %v", err)

		}

		var finalresult domainstructs.Result

		err = json.Unmarshal(httpbody, &finalresult)

		if finalresult != testCase.ExpBody {
			t.Errorf("Expected %#v body, got %#v in case %v", testCase.ExpBody, finalresult, i)
		}

		if finalresult.Is != testCase.ExpResult {
			t.Errorf("Expected %v result; got %v", testCase.ExpResult, finalresult.Is)
		}

		if response.StatusCode != testCase.ExpHttpStatus {
			t.Errorf("Expected %v status code, got %v", testCase.ExpHttpStatus, response.StatusCode)
		}

		if header := response.Header.Get("Content-Type"); header != testCase.ExpHttpHeader {
			t.Errorf("Expected %v header; got %v", testCase.ExpHttpHeader, header)
		}

		err = response.Body.Close()

		if err != nil {
			t.Fatalf("Couldnt close body: %v", err)
		}

	}

}
