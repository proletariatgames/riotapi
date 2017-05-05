package riotapi

import (
	"strconv"
	"testing"
)

func TestLeagueMethod_EntryBySummoner(t *testing.T) {
	var summonerId int64 = 20464313 // YewNewb

	data, err := testClient.League.EntryBySummoner(summonerId, "NA1")

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

	positions, err := testClient.League.PositionsBySummoner(summonerId, "NA1")
	if err != nil {
		t.Errorf("League.PositionsBySummoner error = %v", err)
		return
	}

	if len(positions) == 0 {
		t.Errorf("League.PositionsBySummoner has no leagues; expected at least one league")
		return
	}
}
