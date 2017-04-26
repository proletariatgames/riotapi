package riotapi

import (
	"testing"
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
}
