package riotapi

// Platform struct describes a Riot Platform
type SummonerDTO struct {
	// ID of the summoner icon associated with the summoner.
	ProfileIconId int `json:"profileIconId"`

	// Summoner name.
	Name string `json:"name"`

	// Summoner level associated with the summoner.
	SummonerLevel int64 `json:"summonerLevel"`

	// Date summoner was last modified specified as epoch milliseconds.
	// The following events will update this timestamp: profile icon change,
	// playing the tutorial or advanced tutorial, finishing a game, summoner name change
	RevisionDate int64 `json:"revisionDate"`

	// Summoner ID.
	Id int64 `json:"id"`

	// Account ID.
	AccountId int64 `json:"accountId"`
}

type SummonerMethod struct {
	client *Client
}

func (m *SummonerMethod) ByName(name, platformId string) (*SummonerDTO, error) {
	relPath := "/lol/summoner/v3/summoners/by-name/" + name
	data := new(SummonerDTO)
	if _, err := m.client.get(platformURLBase, relPath, platformId, data); err != nil {
		return nil, err
	}
	return data, nil
}
