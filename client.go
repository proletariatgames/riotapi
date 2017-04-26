package riotapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"gopkg.in/throttled/throttled.v2"
)

var (
	ErrAPIKeyRequired    = errors.New("Required APIKey is missing")
	ErrUnknownPlatformId = errors.New("Unknown platform id")
)

// Version 3+ format, e.g. https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/{name}
const platformURLBase = "https://{platform}.api.riotgames.com"

// Version <3 format, e.g.
const regionURLBase = "https://{region}.api.riotgames.com"

type Client struct {
	client      *http.Client
	APIKey      string
	Summoner    *SummonerMethod
	League      *LeagueMethod
	MatchList   *MatchListMethod
	RateLimiter *throttled.GCRARateLimiter
}

// NewClient initializes and returns a riotapi Client struct
func NewClient(httpClient *http.Client) *Client {
	c := &Client{
		client:      httpClient,
		APIKey:      "",
		RateLimiter: nil,
	}

	c.Summoner = &SummonerMethod{client: c}
	c.League = &LeagueMethod{client: c}
	c.MatchList = &MatchListMethod{client: c}

	return c
}

func (c *Client) get(basePath, relPath, platformId string, decoded interface{}) (*http.Response, error) {
	if len(c.APIKey) == 0 {
		return nil, ErrAPIKeyRequired
	}

	platform := GetPlatform(platformId)
	if platform == nil {
		return nil, ErrUnknownPlatformId
	}

	relURL, err := url.Parse(replaceTokens(relPath, platform))
	if err != nil {
		return nil, err
	}

	baseURL, _ := url.Parse(replaceTokens(basePath, platform))
	combinedURL := baseURL.ResolveReference(relURL)

	// Add the API Key
	q := combinedURL.Query()
	q.Set("api_key", c.APIKey)
	combinedURL.RawQuery = q.Encode()

	// Check rate limiting and wait if necessary
	if c.RateLimiter != nil {
		// rate limits are per region, so use the region in the key
		for {
			limited, context, err := c.RateLimiter.RateLimit("riotapi-"+platform.Id, 1)
			if err != nil {
				return nil, err
			}
			if limited {
				time.Sleep(context.RetryAfter)
			} else {
				break
			}
		}
	}

	req, err := http.NewRequest("GET", combinedURL.String(), nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotModified {
		errBody := new(bytes.Buffer)
		resp.Write(errBody)
		return nil, errors.New("error, status code: " + strconv.Itoa(resp.StatusCode) + ", " + errBody.String())
	}

	if decoded != nil {
		err = json.NewDecoder(resp.Body).Decode(decoded)
	}

	return resp, err
}

func replaceTokens(s string, platform *Platform) string {
	s = strings.Replace(s, "{region}", platform.RegionId, -1)
	s = strings.Replace(s, "{platform}", platform.Id, -1)
	return s
}
