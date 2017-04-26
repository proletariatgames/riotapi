package riotapi

import (
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMatchListMethod_MatchList(t *testing.T) {
	tc := NewClient(new(http.Client))

	tc.APIKey = os.Getenv("RIOTAPI_KEY")

	if len(tc.APIKey) == 0 {
		t.Error("RIOTAPI_KEY environment variable not set. Must be set to execute tests.")
		return
	}

	var summonerId int64 = 20464313 // YewNewb

	options := new(MatchListOptions)
	minusTwoWeeks, err := time.ParseDuration("-336h")
	if err != nil {
		t.Errorf("MatchList.BySummoner could not parse duration")
	}
	options.BeginTime = time.Now().Add(minusTwoWeeks).Unix() * 1000 // seconds to ms
	options.RankedQueues = "TEAM_BUILDER_RANKED_SOLO"

	data, err := tc.MatchList.BySummoner(summonerId, "NA1", options)
	// sleep to stick withing Riot API rate limit of 10/10sec
	time.Sleep(1250 * time.Millisecond)

	if err != nil {
		t.Errorf("MatchList.BySummoner error = %v", err)
		return
	}

	if data == nil {
		t.Error("MatchList.BySummoner returned nil data")
		return
	}
}
