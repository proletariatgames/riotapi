# riotapi
Golang implementation for Riot API

## Status
Very Work In Progress

## APIs Implemented
* LEAGUE-V2.5
* * [Get league entries mapped by summoner ID for a given list of summoner IDs.][league_entries]
* MATCH-V3
* * [Get matchlist for ranked games played on given account ID and platform ID and filtered using given filter parameters, if any.][matchlist]
* SUMMONER-V3
* * [Get a summoner by summoner name.][summoner-byname]

## Rate Limiting
Riot Games has strict API limits for developer keys. As of April 26, 2017 these limits are:
* 10 requests every 10 seconds
* 500 requests every 10 minutes

The `riotapi` package allows the specification of a [Throttled][throttled] rate limiter to limit how frequently
API calls are made to the Riot API. 

*Warning:* by default calls to the `riotapi` package are _not_ limited. If you use this package without
specifying a rate limiter, you may unintentionally exceed the limits allowed by Riot and get your API
access temporarily suspended or permanently banned. Please use caution!

An example rate limiter using the developer limits can be found in [`riotapi_test.go`][test].

## Testing
You will need to set the environment variable `RIOTAPI_KEY` to you Riot Developer API Key in order to run tests
with `go test`. You can obtain your API Key from Riot at [Riot's Developer Website][riotdev].

As mentioned above, the tests specify a rate limiter using the developer limits API rate limits.

[league_entries]: https://developer.riotgames.com/api-methods/#league-v2.5/GET_getLeagueEntriesForSummonerIds
[matchlist]: https://developer.riotgames.com/api-methods/#matchlist-v2.2/GET_getMatchList
[riotdev]: https://developer.riotgames.com
[throttled]: https://github.com/throttled/throttled/
[test]: https://github.com/proletariatgames/riotapi/blob/master/riotapi_test.go
[summoner-byname]: https://developer.riotgames.com/api-methods/#summoner-v3/GET_getBySummonerName