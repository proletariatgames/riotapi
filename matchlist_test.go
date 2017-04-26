package riotapi

import (
	"testing"
	"time"
)

func TestMatchListMethod_MatchList(t *testing.T) {
	var summonerId int64 = 20464313 // YewNewb

	options := new(MatchListOptions)
	minusTwoWeeks, err := time.ParseDuration("-336h")
	if err != nil {
		t.Errorf("MatchList.BySummoner could not parse duration")
	}
	options.BeginTime = time.Now().Add(minusTwoWeeks).Unix() * 1000 // seconds to ms
	options.RankedQueues = "TEAM_BUILDER_RANKED_SOLO"

	data, err := testClient.MatchList.BySummoner(summonerId, "NA1", options)

	if err != nil {
		t.Errorf("MatchList.BySummoner error = %v", err)
		return
	}

	if data == nil {
		t.Error("MatchList.BySummoner returned nil data")
		return
	}
}
