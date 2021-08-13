package muxadapter

import (
	"os"
	"testing"
)

func TestCreateRouter(t *testing.T) {

	obtainedport, router := createRouter()

	muxport := os.Getenv("MuxPort")

	if obtainedport == muxport {

		t.Errorf("Expected port %v, got: %v", muxport, obtainedport)

	}

	if router == nil {
		t.Error()
	}

}
