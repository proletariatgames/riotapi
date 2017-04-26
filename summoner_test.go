package riotapi

import (
	"testing"
	"time"
)

func TestSummonerMethod_ByName(t *testing.T) {
	data, err := testClient.Summoner.ByName("YewNewb", "NA1")

	if err != nil {
		t.Errorf("Summoner.ByName error = %v", err)
		return
	}

	var expectedId int64 = 20464313
	if data.Id != expectedId {
		t.Errorf("Summoner.ByName Id = %v, want %v", data.Id, expectedId)
	}

	// sleep to stick withing Riot API rate limit of 10/10sec
	time.Sleep(1250 * time.Millisecond)
}
