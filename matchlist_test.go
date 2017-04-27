package riotapi

import (
	"testing"
	"time"
)

func TestMatchListMethod_MatchList(t *testing.T) {
	var accountId int64 = 32816596 // YewNewb

	options := new(MatchListOptions)
	minusTwoWeeks, err := time.ParseDuration("-336h")
	if err != nil {
		t.Errorf("MatchList.ByAccount could not parse duration")
	}
	options.BeginTime = time.Now().Add(minusTwoWeeks).Unix() * 1000 // seconds to ms
	options.Queues = make([]QueueType, 2)
	options.Queues[0] = TEAM_BUILDER_RANKED_SOLO
	options.Queues[1] = RANKED_SOLO_5x5

	data, err := testClient.MatchList.ByAccount(accountId, "NA1", options)

	if err != nil {
		t.Errorf("MatchList.ByAccount error = %v", err)
		return
	}

	if data == nil {
		t.Error("MatchList.ByAccount returned nil data")
		return
	}
}
