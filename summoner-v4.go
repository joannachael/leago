package leago

import "context"

type summoner struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	PUUID         string `json:"puuid"`
	Name          string `json:"name"`
	IconID        int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int64  `json:"summonerLevel"`
}

func (c *client) GetSummonerByID(ctx context.Context, region, ID string) (*summoner, error) {
	var summoner summoner
	if err := c.doRequest(ctx, region, "/lol/summoner/v4/summoners/" + ID, &summoner); err != nil {
		return nil, err
	}
	return &summoner, nil
}

func (c *client) GetSummonerByAccountID(ctx context.Context, region, accountID string) (*summoner, error) {
	var summoner summoner
	if err := c.doRequest(ctx, region, "/lol/summoner/v4/summoners/by-account/" + accountID, &summoner); err != nil {
		return nil, err
	}
	return &summoner, nil
}

func (c *client) GetSummonerByName(ctx context.Context, region, name string) (*summoner, error) {
	var summoner summoner
	if err := c.doRequest(ctx, region, "/lol/summoner/v4/summoners/by-name/" + name, &summoner); err != nil {
		return nil, err
	}
	return &summoner, nil
}

func (c *client) GetSummonerByPUUID(ctx context.Context, region, PUUID string) (*summoner, error) {
	var summoner summoner
	if err := c.doRequest(ctx, region, "/lol/summoner/v4/summoners/by-puuid/" + PUUID, &summoner); err != nil {
		return nil, err
	}
	return &summoner, nil
}
