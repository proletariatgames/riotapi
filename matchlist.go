package riotapi

import (
	"strconv"

	"github.com/google/go-querystring/query"
)

// MatchList struct describes a response to a Match List api call
type MatchListDto struct {
	Matches    []MatchReferenceDto `json:"matches,omitempty"`
	TotalGames int                 `json:"totalGames"`
	StartIndex int                 `json:"startIndex"`
	EndIndex   int                 `json:"endIndex"`
}

type MatchReferenceDto struct {
	Lane       string `json:"lane"`
	Champion   int    `json:"champion"`
	PlatformId string `json:"platformId"`
	Timestamp  int64  `json:"timestamp"` // Epoch milliseconds
	Region     string `json:"region"`
	GameId     int64  `json:"gameId"`
	Queue      int    `json:"queue"`
	Role       string `json:"role"`
	Season     int    `json:"season"`
}

type MatchListOptions struct {
	// The end time to use for fetching games specified as epoch milliseconds.
	EndTime int64 `url:"endTime,omitempty"`

	// The end index to use for fetching games.
	EndIndex int `url:"endIndex,omitempty"`

	// Set of queue IDs for filtering matchlist.
	Queues []QueueType `url:"queue,omitempty"`

	// Set of season IDs for filtering matchlist.
	Seasons []int `url:"season,omitempty"`

	// The begin time to use for fetching games specified as epoch milliseconds.
	BeginTime int64 `url:"beginTime,omitempty"`

	// The begin index to use for fetching games.
	BeginIndex int `url:"beginIndex,omitempty"`

	// Comma-separated list of champion IDs to use for fetching games.
	Champions []int `url:"champion,omitempty"`
}

type MatchListMethod struct {
	client *Client
}

func (m *MatchListMethod) ByAccount(accountId int64, platformId string, options *MatchListOptions) (*MatchListDto, error) {
	relPath := "/lol/match/v3/matchlists/by-account/" + strconv.FormatInt(accountId, 10)
	if options != nil {
		if vals, err := query.Values(options); err != nil {
			return nil, err
		} else {
			relPath += "?" + vals.Encode()
		}
	}
	data := new(MatchListDto)
	if _, err := m.client.get(platformURLBase, relPath, platformId, data); err != nil {
		return nil, err
	}
	return data, nil
}
