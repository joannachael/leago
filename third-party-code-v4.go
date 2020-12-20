package leago

import "context"

func (c *client) GetPartyCodeBySummonerId(ctx context.Context, region, Id string) (string, error) {
	var code string
	if err := c.doRequest(ctx, region, "/lol/platform/v4/third-party-code/by-summoner/" + Id, &code); err != nil {
		return "", err
	}
	return code, nil
}
