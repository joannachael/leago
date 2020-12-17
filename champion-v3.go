package leago

import "context"

type rotation struct {
	FreeChampionsIDs               []int `json:"freeChampionIds"`
	FreeChampionsIDsForNewsPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel              int   `json:"maxNewPlayerLevel"`
}

func (c *client) GetRotations(ctx context.Context, region string) (*rotation, error) {
	var rotation rotation
	if err := c.doRequest(ctx, region, "/lol/platform/v3/champion-rotations/", &rotation); err != nil {
		return nil, err
	}
	return &rotation, nil
}
