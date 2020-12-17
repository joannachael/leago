package leago

import "context"

func (c *client) GetPartyCodeBySummonerID(ctx context.Context, region, ID string) (string, error) {
	var code string
	if err := c.doRequest(ctx, region, "/lol/platform/v4/third-party-code/by-summoner/" + ID, &code); err != nil {
		return "", err
	}
	return code, nil
}
