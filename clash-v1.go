package leago

import (
	"context"
	"strconv"
)

type (
	team struct {
		ID           string        `json:"id"`
		TournamentID int           `json:"tournamentId"`
		Name         string        `json:"name"`
		IconID       int           `json:"iconId"`
		Tier         int           `json:"tier"`
		Captain      string        `json:"captain"`
		Abbreviation string        `json:"abbreviation"`
		Players      []clashPlayer `json:"players"`
	}

	clashPlayer struct {
		SummonerID string `json:"summonerId"`
		TeamID     string `json:"teamId"`
		Position   string `json:"position"`
		Role       string `json:"role"`
	}

	tournament struct {
		ID               int     `json:"id"`
		ThemeID          int     `json:"themeId"`
		NameKey          string  `json:"nameKey"`
		NameKeySecondary string  `json:"nameKeySecondary"`
		Schedule         []phase `json:"schedule"`
	}

	phase struct {
		ID               int   `json:"id"`
		RegistrationTime int64 `json:"registrationTime"`
		StartTime        int64 `json:"startTime"`
		Cancelled        bool  `json:"cancelled"`
	}
)

func (c *client) GetPlayersBySummonerID(ctx context.Context, region, ID string) ([]clashPlayer, error) {
	var players []clashPlayer
	if err := c.doRequest(ctx, region, "/lol/clash/v1/players/by-summoner/" + ID, &players); err != nil {
		return nil, err
	}
	return players, nil
}

func (c *client) GetTeamByID(ctx context.Context, region, ID string) (*team, error) {
	var team team
	if err := c.doRequest(ctx, region, "/lol/clash/v1/teams/" + ID, &team); err != nil {
		return nil, err
	}
	return &team, nil
}

func (c *client) GetTournaments(ctx context.Context, region string) ([]tournament, error) {
	var tournaments []tournament
	if err := c.doRequest(ctx, region, "/lol/clash/v1/tournaments", &tournaments); err != nil {
		return nil, err
	}
	return tournaments, nil
}

func (c *client) GetTournamentByTeamID(ctx context.Context, region, ID string) (*tournament, error) {
	var tournament tournament
	if err := c.doRequest(ctx, region, "/lol/clash/v1/tournaments/by-team/" + ID, &tournament); err != nil {
		return nil, err
	}
	return &tournament, nil
}

func (c *client) GetTournamentByID(ctx context.Context, region string, ID int) (*tournament, error) {
	var tournament tournament
	if err := c.doRequest(ctx, region, "/lol/clash/v1/tournaments/" + strconv.Itoa(ID), &tournament); err != nil {
		return nil, err
	}
	return &tournament, nil
}
