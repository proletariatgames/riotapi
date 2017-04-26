package riotapi

import (
	"net/http"
	"os"
	"testing"
)

// Client shared by all tests
var testClient *Client

func TestMain(m *testing.M) {
	testClient = NewClient(new(http.Client))

	testClient.APIKey = os.Getenv("RIOTAPI_KEY")

	if len(testClient.APIKey) == 0 {
		println("RIOTAPI_KEY environment variable not set. Must be set to execute tests.")
		os.Exit(1)
	}

	os.Exit(m.Run())
}
