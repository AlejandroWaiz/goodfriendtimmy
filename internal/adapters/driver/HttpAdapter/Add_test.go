package muxadapter

import (
	"bytes"
	"encoding/json"
	domainstructs "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/Structs"
	timmyadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/TimmyAdapter"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdd(t *testing.T) {

	testTable := []struct {
		FirstOperand   float64
		SecondOperand  float64
		ExpectedResult float64
	}{
		{FirstOperand: 1, SecondOperand: 2, ExpectedResult: 3},
	}

	for _, testCase := range testTable {

		opbody := domainstructs.Operation{FirstOperand: testCase.FirstOperand, SecondOperand: testCase.SecondOperand}

		bodyInBytes, err := json.Marshal(opbody)

		if err != nil {
			t.Log(err)
		}

		buf := bytes.NewBuffer(bodyInBytes)

		if err != nil {
			t.Log(err)
		}

		request, err := http.NewRequest("POST", "http://localhost:8080/Add", buf)

		if err != nil {
			t.Log(err)
		}

		recorder := httptest.NewRecorder()

		timmy := timmyadapter.CreateTimmyAdapter()
		muxAdapter := MuxAdapter{domainport: timmy}

		func() {

			muxAdapter.AddHttpHandler(recorder, request)

			response := recorder.Result()
			defer response.Body.Close()

			httpbody, err := ioutil.ReadAll(response.Body)
			log.Println(httpbody)

			if err != nil {

				t.Fatalf("Could not read response: %v", err)

			}

			var finalresult domainstructs.Result

			err = json.Unmarshal(httpbody, &finalresult)

			log.Println(finalresult)

			/*err = json.NewDecoder(response.Body).Decode(&finalresult)
			log.Println(finalresult)

			if err != nil {
				t.Log(err)
				t.Fatalf("Bad format. want %T got %T", response.Body ,finalresult)
			}*/

			if response.StatusCode != http.StatusOK {

				t.Errorf("Expected status 200 (OK), got %v", response.StatusCode)

			}

			//Here we test the double of given value. Change this if you change the default value. Default = 10.
			if finalresult.Is != testCase.ExpectedResult {

				t.Fatalf("Expected %v; got %v", testCase.ExpectedResult, finalresult.Is)

			}

		}()

	}

}
