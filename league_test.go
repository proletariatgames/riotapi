package riotapi

import (
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestLeagueMethod_EntryBySummoner(t *testing.T) {
	tc := NewClient(new(http.Client))

	tc.APIKey = os.Getenv("RIOTAPI_KEY")

	if len(tc.APIKey) == 0 {
		t.Error("RIOTAPI_KEY environment variable not set. Must be set to execute tests.")
		return
	}

	var summonerId int64 = 20464313 // YewNewb

	data, err := tc.League.EntryBySummoner(summonerId, "NA1")
	// sleep to stick withing Riot API rate limit of 10/10sec
	time.Sleep(1250 * time.Millisecond)

	if err != nil {
		t.Errorf("League.EntryBySummoner error = %v", err)
		return
	}

	summonerIdStr := strconv.FormatInt(summonerId, 10)
	leagues := data[summonerIdStr]
	if leagues == nil {
		t.Errorf("League.EntryBySummoner summonerId was not found; expected = %v", summonerIdStr)
		return
	}

	if len(leagues) == 0 {
		t.Errorf("League.EntryBySummoner has no leagues; expected at least one league")
		return
	}

	if len(leagues[0].Entries) == 0 {
		t.Errorf("League.EntryBySummoner first league (queue = %v, tier = %v) has no entries; expected at least one entry", leagues[0].Queue, leagues[0].Tier)
	}
}
