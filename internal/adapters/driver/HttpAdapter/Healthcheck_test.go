package muxadapter

import (
	"net/http"
	"net/http/httptest"
	"testing"

	timmyadapter "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain/TimmyAdapter"
)

func TestHealthcheck(t *testing.T) {

	recorder := httptest.NewRecorder()

	//errReader declared in Common_test.go file
	request, err := http.NewRequest("GET", "http://localhost:8080/Healthcheck", nil)

	if err != nil {
		t.Log(err)
	}

	timmy := timmyadapter.CreateTimmyAdapter()
	muxAdapter := MuxAdapter{domainport: timmy}

	muxAdapter.HealthCheck(recorder, request)

}
