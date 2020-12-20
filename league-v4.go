package leago

import (
	"context"
	"strconv"
)

type (
	miniSeries struct {
		Losses   int    `json:"losses"`
		Progress string `json:"progress"`
		Target   int    `json:"target"`
		Wins     int    `json:"wins"`
	}

	item struct {
		FreshBlood   bool       `json:"freshBlood"`
		Wins         int        `json:"wins"`
		SummonerName string     `json:"summonerName"`
		MiniSeries   miniSeries `json:"miniSeries"`
		Inactive     bool       `json:"inactive"`
		Veteran      bool       `json:"veteran"`
		HotStreak    bool       `json:"hotStreak"`
		Rank         string     `json:"rank"`
		Points       int        `json:"leaguePoints"`
		Losses       int        `json:"losses"`
		SummonerId   string     `json:"summonerId"`
	}

	entry struct {
		Id           string     `json:"leagueId"`
		SummonerId   string     `json:"summonerId"`
		SummonerName string     `json:"summonerName"`
		QueueType    string     `json:"queueType"`
		Tier         string     `json:"tier"`
		Rank         string     `json:"rank"`
		Points       int        `json:"leaguePoints"`
		Wins         int        `json:"wins"`
		Losses       int        `json:"losses"`
		HotStreak    bool       `json:"hotStreak"`
		Veteran      bool       `json:"veteran"`
		FreshBlood   bool       `json:"freshBlood"`
		Inactive     bool       `json:"inactive"`
		MiniSeries   miniSeries `json:"miniSeries"`
	}

	leagueList struct {
		Id      string `json:"leagueId"`
		Name    string `json:"name"`
		Entries []item `json:"entries"`
		Tier    string `json:"tier"`
		Queue   string `json:"queue"`
	}
)

func (c *client) GetChallengerLeagueByQueue(ctx context.Context, region, queue string) (*leagueList, error) {
	var leagueList leagueList
	if err := c.doRequest(ctx, region, "/lol/league/v4/challengerleagues/by-queue/" + queue, &leagueList); err != nil {
		return nil, err
	}
	return &leagueList, nil
}

func (c *client) GetEntriesBySummonerId(ctx context.Context, region, summonerId string) ([]entry, error) {
	var entries []entry
	if err := c.doRequest(ctx, region, "/lol/league/v4/entries/by-summoner/" + summonerId, &entries); err != nil {
		return nil, err
	}
	return entries, nil
}

func (c *client) GetEntriesByQueueTierAndDivision(ctx context.Context, region, queue, tier, division string, page int) ([]entry, error) {
	var entries []entry
	address := "/lol/league/v4/entries/" + queue + "/" + tier + "/" + division + "?page=" + strconv.Itoa(page)
	if err := c.doRequest(ctx, region, address, &entries); err != nil {
		return nil, err
	}
	return entries, nil
}

func (c *client) GetGrandmasterLeagueByQueue(ctx context.Context, region, queue string) (*leagueList, error) {
	var leagueList leagueList
	if err := c.doRequest(ctx, region, "/lol/league/v4/grandmasterleagues/by-queue/" + queue, &leagueList); err != nil {
		return nil, err
	}
	return &leagueList, nil
}

func (c *client) GetLeagueById(ctx context.Context, region, leagueId string) (*leagueList, error) {
	var leagueList leagueList
	if err := c.doRequest(ctx, region, "/lol/league/v4/leagues/" + leagueId, &leagueList); err != nil {
		return nil, err
	}
	return &leagueList, nil
}

func (c *client) GetMasterLeagueByQueue(ctx context.Context, region, queue string) (*leagueList, error) {
	var leagueList leagueList
	if err := c.doRequest(ctx, region, "/lol/league/v4/masterleagues/by-queue/" + queue, &leagueList); err != nil {
		return nil, err
	}
	return &leagueList, nil
}
