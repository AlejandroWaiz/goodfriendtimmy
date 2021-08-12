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

func TestDivide(t *testing.T) {

	testTable := []struct {
		FirstOperand  float64 `json:",omitempty"`
		SecondOperand float64 `json:",omitempty"`
		ExpResult     float64
		ExpHttpStatus int
		ExpHttpHeader string
		ExpBody       domainstructs.Result
		ExpErr        error
	}{
		{FirstOperand: 5, SecondOperand: 5, ExpResult: 1, ExpHttpStatus: 200, ExpHttpHeader: "application/json", ExpBody: domainstructs.Result{Is: 1}, ExpErr: nil},
		{FirstOperand: 1, SecondOperand: 0, ExpResult: 0, ExpHttpStatus: 400, ExpHttpHeader: "application/json", ExpBody: domainstructs.Result{Is: 0}, ExpErr: &timmyadapter.DivideByZero{}},
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

		request, err := http.NewRequest("POST", "http://localhost:8080/Divide", buf)
		request.Header.Set("Content-type", "application/json")

		if err != nil {
			t.Log(err)
		}

		recorder := httptest.NewRecorder()

		timmy := timmyadapter.CreateTimmyAdapter()
		muxAdapter := MuxAdapter{domainport: timmy}

		muxAdapter.DivideHttpHandler(recorder, request)

		response := recorder.Result()

		httpbody, err := ioutil.ReadAll(response.Body)

		if err != nil {

			t.Fatalf("Could not read response: %v", err)

		}

		var finalresult domainstructs.Result

		err = json.Unmarshal(httpbody, &finalresult)

		if finalresult != testCase.ExpBody {
			t.Errorf("Expected %#v on body, got %#v in case %v", testCase.ExpBody, finalresult, i)
		}

		if finalresult.Is != testCase.ExpResult {
			t.Errorf("Expected %v on result; got %v in case %v", testCase.ExpResult, finalresult.Is, i)
		}

		if response.StatusCode != testCase.ExpHttpStatus {
			t.Errorf("Expected %v on status code, got %v in case %v", testCase.ExpHttpStatus, response.StatusCode, i)
		}

		if header := response.Header.Get("Content-Type"); header != testCase.ExpHttpHeader {
			t.Errorf("Expected %v on header; got %v in case %v", testCase.ExpHttpHeader, header, i)
		}

		err = response.Body.Close()

		if err != nil {
			t.Fatalf("Couldnt close body: %v in case %v", err, i)
		}

	}

}
