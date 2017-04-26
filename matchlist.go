package riotapi

import (
	"strconv"

	"github.com/google/go-querystring/query"
)

// MatchList struct describes a response to a Match List api call
type MatchList struct {
	Matches    []MatchReference `json:"matches,omitempty"`
	TotalGames int              `json:"totalGames"`
	StartIndex int              `json:"startIndex"`
	EndIndex   int              `json:"endIndex"`
}

type MatchReference struct {
	Lane       string `json:"lane"`
	Champion   int64  `json:"champion"`
	PlatformId string `json:"platformId"`
	Timestamp  int64  `json:"timestamp"` // Epoch milliseconds
	Region     string `json:"region"`
	MatchId    int64  `json:"matchId"`
	Queue      string `json:"queue"`
	Role       string `json:"role"`
	Season     string `json:"season"`
}

type MatchListOptions struct {
	// The end time to use for fetching games specified as epoch milliseconds.
	EndTime int64 `url:"endTime,omitempty"`

	// The end index to use for fetching games.
	EndIndex int `url:"endIndex,omitempty"`

	// Comma-separated list of ranked queue types to use for fetching games.
	// Non-ranked queue types will be ignored.
	// Values:
	//   RANKED_FLEX_SR
	//   RANKED_SOLO_5x5
	//   RANKED_TEAM_3x3
	//   RANKED_TEAM_5x5
	//   TEAM_BUILDER_DRAFT_RANKED_5x5
	//   TEAM_BUILDER_RANKED_SOLO
	RankedQueues string `url:"rankedQueues,omitempty"`

	// Comma-separated list of seasons to use for fetching games.
	// Values:
	//   PRESEASON3
	//   SEASON3
	//   PRESEASON2014
	//   SEASON2014
	//   PRESEASON2015
	//   SEASON2015
	//   PRESEASON2016
	//   SEASON2016
	//   PRESEASON2017
	//   SEASON2017
	Seasons string `url:"seasons,omitempty"`

	// The begin time to use for fetching games specified as epoch milliseconds.
	BeginTime int64 `url:"beginTime,omitempty"`

	// The begin index to use for fetching games.
	BeginIndex int `url:"beginIndex,omitempty"`

	// Comma-separated list of champion IDs to use for fetching games.
	ChampionIds string `url:"championIds,omitempty"`
}

type MatchListMethod struct {
	client *Client
}

func (m *MatchListMethod) BySummoner(summonerId int64, platformId string, options *MatchListOptions) (*MatchList, error) {
	relPath := "/api/lol/{region}/v2.2/matchlist/by-summoner/" + strconv.FormatInt(summonerId, 10)
	if options != nil {
		if vals, err := query.Values(options); err != nil {
			return nil, err
		} else {
			relPath += "?" + vals.Encode()
		}
	}
	data := new(MatchList)
	if _, err := m.client.get(regionURLBase, relPath, platformId, data); err != nil {
		return nil, err
	}
	return data, nil
}
