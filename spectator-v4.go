package leago

import "context"

type (
	currentGame struct {
		GameId            int64                    `json:"gameId"`
		GameType          string                   `json:"gameType"`
		GameStartTime     int64                    `json:"gameStartTime"`
		MapId             int64                    `json:"mapId"`
		GameLength        int64                    `json:"gameLength"`
		PlatformId        string                   `json:"platformId"`
		GameMode          string                   `json:"gameMode"`
		BannedChampions   []bannedChampion         `json:"bannedChampions"`
		GameQueueConfigId int64                    `json:"gameQueueConfigId"`
		Observers         observer                 `json:"observers"`
		Participants      []currentGameParticipant `json:"participants"`
	}

	bannedChampion struct {
		PickTurn   int   `json:"pickTurn"`
		ChampionId int64 `json:"championId"`
		TeamId     int64 `json:"teamId"`
	}

	observer struct {
		EncryptionKey string `json:"encryptionKey"`
	}

	currentGameParticipant struct {
		ChampionId               int64                 `json:"championId"`
		Perks                    perks                 `json:"perks"`
		ProfileIconId            int64                 `json:"profileIconId"`
		Bot                      bool                  `json:"bot"`
		TeamId                   int64                 `json:"teamId"`
		SummonerName             string                `json:"summonerName"`
		SummonerId               string                `json:"summonerId"`
		Spell1Id                 int64                 `json:"spell1Id"`
		Spell2Id                 int64                 `json:"spell2Id"`
		GameCustomizationObjects []customizationObject `json:"gameCustomizationObjects"`
	}

	perks struct {
		PerkIds      []int64 `json:"perkIds"`
		PerkStyle    int64   `json:"perkStyle"`
		PerkSubStyle int64   `json:"perkSubStyle"`
	}

	customizationObject struct {
		Category string `json:"category"`
		Content  string `json:"content"`
	}

	featuredGame struct {
		GameMode          string                    `json:"gameMode"`
		GameLength        int64                     `json:"gameLength"`
		MapId             int64                     `json:"mapId"`
		GameType          string                    `json:"gameType"`
		BannedChampions   []bannedChampion          `json:"bannedChampions"`
		GameId            int64                     `json:"gameId"`
		Observers         observer                  `json:"observers"`
		GameQueueConfigId int64                     `json:"gameQueueConfigId"`
		GameStartTime     int64                     `json:"gameStartTime"`
		Participants      []featuredGameParticipant `json:"participants"`
		PlatformId        string                    `json:"platformId"`
	}

	featuredGameParticipant struct {
		Bot           bool   `json:"bot"`
		Spell1Id      int64  `json:"spell1Id"`
		Spell2Id      int64  `json:"spell2Id"`
		ProfileIconId int64  `json:"profileIconId"`
		SummonerName  string `json:"summonerName"`
		ChampionId    int64  `json:"championId"`
		TeamId        int64  `json:"teamId"`
	}

	featuredGames struct {
		GameList              []featuredGame `json:"gameList"`
		ClientRefreshInterval int64          `json:"clientRefreshInterval"`
	}
)

func (c *client) GetCurrentGameBySummonerId(ctx context.Context, region, summonerId string) (*currentGame, error) {
	var currentGame currentGame
	if err := c.doRequest(ctx, region, "/lol/spectator/v4/active-games/by-summoner/" + summonerId, &currentGame); err != nil {
		return nil, err
	}
	return &currentGame, nil
}

func (c *client) GetFeaturedGames(ctx context.Context, region string) (*featuredGames, error) {
	var featuredGames featuredGames
	if err := c.doRequest(ctx, region, "/lol/spectator/v4/featured-games/", &featuredGames); err != nil {
		return nil, err
	}
	return &featuredGames, nil
}
