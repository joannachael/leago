package leago

import (
	"context"
	"strconv"
)

type championMastery struct {
	SummonerID                   string `json:"summonerId"`
	ChampionID                   int64  `json:"championId"`
	ChampionLevel                int    `json:"championLevel"`
	ChampionPoints               int    `json:"championPoints"`
	LastPlayTime                 int64  `json:"lastPlayTime"`
	ChampionPointsSinceLastLevel int64  `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int64  `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	TokensEarned                 int    `json:"tokensEarned"`
}

func (c *client) GetMasteriesBySummonerID(ctx context.Context, region, ID string) ([]championMastery, error) {
	var masteries []championMastery
	address := "/lol/champion-mastery/v4/champion-masteries/by-summoner/"
	if err := c.doRequest(ctx, region, address + ID, &masteries); err != nil {
		return nil, err
	}
	return masteries, nil
}

func (c *client) GetMasteryBySummonerAndChampionID(
	ctx context.Context, region, summonerID string, championID int,
) (*championMastery, error) {
	var mastery championMastery
	api := "/lol/champion-mastery/v4/champion-masteries/by-summoner/"
	address := api + summonerID + "/by-champion/" + strconv.Itoa(championID)
	if err := c.doRequest(ctx, region, address, &mastery); err != nil {
		return nil, err
	}
	return &mastery, nil
}

func (c *client) GetScoreBySummonerID(ctx context.Context, region, ID string) (int, error) {
	var score int
	if err := c.doRequest(ctx, region, "/lol/champion-mastery/v4/scores/by-summoner/" + ID, &score); err != nil {
		return 0, err
	}
	return score, nil
}
