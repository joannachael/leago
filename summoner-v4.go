package leago

import "context"

type summoner struct {
	Id            string `json:"id"`
	AccountId     string `json:"accountId"`
	PUUID         string `json:"puuid"`
	Name          string `json:"name"`
	IconId        int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int64  `json:"summonerLevel"`
}

func (c *client) GetSummonerById(ctx context.Context, region, summonerId string) (*summoner, error) {
	var summoner summoner
	if err := c.doRequest(ctx, region, "/lol/summoner/v4/summoners/" + summonerId, &summoner); err != nil {
		return nil, err
	}
	return &summoner, nil
}

func (c *client) GetSummonerByAccountId(ctx context.Context, region, accountId string) (*summoner, error) {
	var summoner summoner
	if err := c.doRequest(ctx, region, "/lol/summoner/v4/summoners/by-account/" + accountId, &summoner); err != nil {
		return nil, err
	}
	return &summoner, nil
}

func (c *client) GetSummonerByName(ctx context.Context, region, summonerName string) (*summoner, error) {
	var summoner summoner
	if err := c.doRequest(ctx, region, "/lol/summoner/v4/summoners/by-name/" + summonerName, &summoner); err != nil {
		return nil, err
	}
	return &summoner, nil
}

func (c *client) GetSummonerByPUUID(ctx context.Context, region, summonerPUUID string) (*summoner, error) {
	var summoner summoner
	if err := c.doRequest(ctx, region, "/lol/summoner/v4/summoners/by-puuid/" + summonerPUUID, &summoner); err != nil {
		return nil, err
	}
	return &summoner, nil
}
