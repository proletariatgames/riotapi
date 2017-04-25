package riotapi

import (
	"net/http"
	"os"
	"testing"
	"time"
)

func TestSummonerMethod_SummonersByName(t *testing.T) {
	tc := NewClient(new(http.Client))

	tc.APIKey = os.Getenv("RIOTAPI_KEY")

	if len(tc.APIKey) == 0 {
		t.Error("RIOTAPI_KEY environment variable not set. Must be set to execute tests.")
	}
	data, err := tc.Summoner.SummonersByName("YewNewb", "NA1")

	if err != nil {
		t.Errorf("Summoner.SummonersByName error = %v", err)
		return
	}

	var expectedId int64 = 20464313
	if data.Id != expectedId {
		t.Errorf("Summoner.SummonersByName Id = %v, want %v", data.Id, expectedId)
	}

	// sleep to stick withing Riot API rate limit of 10/10sec
	time.Sleep(1250 * time.Millisecond)
}
