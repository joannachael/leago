package leago

import "context"

func (c *client) GetPartyCodeBySummonerId(ctx context.Context, region, summonerId string) (string, error) {
	var code string
	if err := c.doRequest(ctx, region, "/lol/platform/v4/third-party-code/by-summoner/" + summonerId, &code); err != nil {
		return "", err
	}
	return code, nil
}
