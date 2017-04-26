package riotapi

import (
	"strconv"
)

// Map of summonerId (string version) to slice of LeagueDTO structures
type LeagueResponseMap map[string][]LeagueDTO

// LeagueDTO struct describes a response to a League api call
type LeagueDTO struct {
	// The league's queue type.
	Queue string `json:"queue"`

	// The league's tier.
	Tier string `json:"tier"`

	// Specifies the relevant participant that is a member of this league
	// (i.e., a requested summoner ID, a requested team ID, or the ID of
	// a team to which one of the requested summoners belongs). Only
	// present when full league is requested so that participant's entry
	// can be identified. Not present when individual entry is requested.
	ParticipantId string `json:"participantId,omitempty"`

	// This name is an internal place-holder name only. Display and
	// localization of names in the game client are handled client-side.
	Name string `json:"name,omitempty"`

	// The requested league entries.
	Entries []LeagueEntryDTO `json:"entries"`
}

// LeagueEntryDTO describes an entry in response to a League api call
type LeagueEntryDTO struct {
	// Specifies if the participant is fresh blood.
	IsFreshBlood bool `json:"isFreshBlood"`

	// The league division of the participant.
	Division string `json:"division"`

	// The playstyle of the participant.
	Playstyle string `json:"playstyle"`

	// Mini series data for the participant. Only present if the participant
	// is currently in a mini series.
	MiniSeries *MiniSeriesDTO `json:"miniSeries,omitempty"`

	// The number of wins for the participant.
	Wins int `json:"wins"`

	// The number of losses for the participant.
	Losses int `json:"losses"`

	// The ID of the participant (i.e., summoner or team) represented by
	// this entry.
	PlayerOrTeamId string `json:"playerOrTeamId"`

	// The name of the the participant (i.e., summoner or team) represented
	// by this entry.
	PlayerOrTeamName string `json:"playerOrTeamName"`

	//  Specifies if the participant is on a hot streak.
	IsHotStreak bool `json:"isHotStreak"`

	//  Specifies if the participant is inactive.
	IsInactive bool `json:"isInactive"`

	// Specifies if the participant is a veteran.
	IsVeteran bool `json:"isVeteran"`

	// The league points of the participant.
	LeaguePoints int `json:"leaguePoints"`
}

type MiniSeriesDTO struct {
	// String showing the current, sequential mini series progress where
	// 'W' represents a win, 'L' represents a loss, and 'N' represents a
	// game that hasn't been played yet.
	Progress string `json:"progress"`

	// Number of current losses in the mini series.
	Losses int `json:"losses"`

	// Number of wins required for promotion.
	Target int `json:"target"`

	// Number of current wins in the mini series.
	Wins int `json:"wins"`
}

type LeagueMethod struct {
	client *Client
}

func (m *LeagueMethod) EntryBySummoner(summonerId int64, platformId string) (LeagueResponseMap, error) {
	relPath := "/api/lol/{region}/v2.5/league/by-summoner/" + strconv.FormatInt(summonerId, 10) + "/entry"
	data := make(LeagueResponseMap)
	if _, err := m.client.get(regionURLBase, relPath, platformId, &data); err != nil {
		return nil, err
	}
	return data, nil
}
