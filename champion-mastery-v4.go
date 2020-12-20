package leago

import (
	"context"
	"strconv"
)

type championMastery struct {
	SummonerId                   string `json:"summonerId"`
	ChampionId                   int64  `json:"championId"`
	ChampionLevel                int    `json:"championLevel"`
	ChampionPoints               int    `json:"championPoints"`
	LastPlayTime                 int64  `json:"lastPlayTime"`
	ChampionPointsSinceLastLevel int64  `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int64  `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	TokensEarned                 int    `json:"tokensEarned"`
}

func (c *client) GetMasteriesBySummonerId(ctx context.Context, region, summonerId string) ([]championMastery, error) {
	var masteries []championMastery
	address := "/lol/champion-mastery/v4/champion-masteries/by-summoner/" + summonerId
	if err := c.doRequest(ctx, region, address, &masteries); err != nil {
		return nil, err
	}
	return masteries, nil
}

func (c *client) GetMasteryBySummonerAndChampionId(ctx context.Context, region, summonerId string, championId int) (*championMastery, error) {
	var mastery championMastery
	api := "/lol/champion-mastery/v4/champion-masteries/by-summoner/"
	address := api + summonerId + "/by-champion/" + strconv.Itoa(championId)
	if err := c.doRequest(ctx, region, address, &mastery); err != nil {
		return nil, err
	}
	return &mastery, nil
}

func (c *client) GetScoreBySummonerId(ctx context.Context, region, summonerId string) (int, error) {
	var score int
	if err := c.doRequest(ctx, region, "/lol/champion-mastery/v4/scores/by-summoner/" + summonerId, &score); err != nil {
		return 0, err
	}
	return score, nil
}
