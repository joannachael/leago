package leago

import (
	"context"
	"strconv"
)

type (
	team struct {
		Id           string        `json:"id"`
		TournamentId int           `json:"tournamentId"`
		Name         string        `json:"name"`
		IconId       int           `json:"iconId"`
		Tier         int           `json:"tier"`
		Captain      string        `json:"captain"`
		Abbreviation string        `json:"abbreviation"`
		Players      []clashPlayer `json:"players"`
	}

	clashPlayer struct {
		SummonerId string `json:"summonerId"`
		TeamId     string `json:"teamId"`
		Position   string `json:"position"`
		Role       string `json:"role"`
	}

	tournament struct {
		Id               int     `json:"id"`
		ThemeId          int     `json:"themeId"`
		NameKey          string  `json:"nameKey"`
		NameKeySecondary string  `json:"nameKeySecondary"`
		Schedule         []phase `json:"schedule"`
	}

	phase struct {
		Id               int   `json:"id"`
		RegistrationTime int64 `json:"registrationTime"`
		StartTime        int64 `json:"startTime"`
		Cancelled        bool  `json:"cancelled"`
	}
)

func (c *client) GetPlayersBySummonerId(ctx context.Context, region, summonerId string) ([]clashPlayer, error) {
	var players []clashPlayer
	if err := c.doRequest(ctx, region, "/lol/clash/v1/players/by-summoner/" + summonerId, &players); err != nil {
		return nil, err
	}
	return players, nil
}

func (c *client) GetTeamById(ctx context.Context, region, teamId string) (*team, error) {
	var team team
	if err := c.doRequest(ctx, region, "/lol/clash/v1/teams/" + teamId, &team); err != nil {
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

func (c *client) GetTournamentByTeamId(ctx context.Context, region, teamId string) (*tournament, error) {
	var tournament tournament
	if err := c.doRequest(ctx, region, "/lol/clash/v1/tournaments/by-team/" + teamId, &tournament); err != nil {
		return nil, err
	}
	return &tournament, nil
}

func (c *client) GetTournamentById(ctx context.Context, region string, tournamentId int) (*tournament, error) {
	var tournament tournament
	if err := c.doRequest(ctx, region, "/lol/clash/v1/tournaments/" + strconv.Itoa(tournamentId), &tournament); err != nil {
		return nil, err
	}
	return &tournament, nil
}
