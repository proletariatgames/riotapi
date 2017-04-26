package riotapi

import (
	"net/http"
	"os"
	"testing"

	"gopkg.in/throttled/throttled.v2"
	memstore "gopkg.in/throttled/throttled.v2/store/memstore"
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

	// Set up rate limiting to 50/minute to fit the 500/10 minutes developer
	// limit

	// Data store. Since we only vary keys by platformId, use that to calculate size
	store, err := memstore.New(1024 * len(Platforms))
	if err != nil {
		println("Unable to create rate limiter datastore, err: ", err.Error())
		os.Exit(2)
	}

	quota := throttled.RateQuota{throttled.PerMin(50), 0}
	testClient.RateLimiter, err = throttled.NewGCRARateLimiter(store, quota)
	if err != nil {
		println("Unable to create rate limiter, err: ", err.Error())
		os.Exit(3)
	}

	os.Exit(m.Run())
}
