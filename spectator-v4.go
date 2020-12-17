package leago

import "context"

type (
	currentGame struct {
		GameID            int64                    `json:"gameId"`
		GameType          string                   `json:"gameType"`
		GameStartTime     int64                    `json:"gameStartTime"`
		MapID             int64                    `json:"mapId"`
		GameLength        int64                    `json:"gameLength"`
		PlatformID        string                   `json:"platformId"`
		GameMode          string                   `json:"gameMode"`
		BannedChampions   []bannedChampion         `json:"bannedChampions"`
		GameQueueConfigID int64                    `json:"gameQueueConfigId"`
		Observers         observer                 `json:"observers"`
		Participants      []currentGameParticipant `json:"participants"`
	}

	bannedChampion struct {
		PickTurn   int   `json:"pickTurn"`
		ChampionID int64 `json:"championId"`
		TeamID     int64 `json:"teamId"`
	}

	observer struct {
		EncryptionKey string `json:"encryptionKey"`
	}

	currentGameParticipant struct {
		ChampionID               int64                 `json:"championId"`
		Perks                    perks                 `json:"perks"`
		ProfileIconID            int64                 `json:"profileIconId"`
		Bot                      bool                  `json:"bot"`
		TeamId                   int64                 `json:"teamId"`
		SummonerName             string                `json:"summonerName"`
		SummonerID               string                `json:"summonerId"`
		Spell1ID                 int64                 `json:"spell1Id"`
		Spell2ID                 int64                 `json:"spell2Id"`
		GameCustomizationObjects []customizationObject `json:"gameCustomizationObjects"`
	}

	perks struct {
		PerkIDs      []int64 `json:"perkIds"`
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
		MapID             int64                     `json:"mapId"`
		GameType          string                    `json:"gameType"`
		BannedChampions   []bannedChampion          `json:"bannedChampions"`
		GameID            int64                     `json:"gameId"`
		Observers         observer                  `json:"observers"`
		GameQueueConfigID int64                     `json:"gameQueueConfigId"`
		GameStartTime     int64                     `json:"gameStartTime"`
		Participants      []featuredGameParticipant `json:"participants"`
		PlatformID        string                    `json:"platformId"`
	}

	featuredGameParticipant struct {
		Bot           bool   `json:"bot"`
		Spell1ID      int64  `json:"spell1Id"`
		Spell2ID      int64  `json:"spell2Id"`
		ProfileIconID int64  `json:"profileIconId"`
		SummonerName  string `json:"summonerName"`
		ChampionID    int64  `json:"championId"`
		TeamID        int64  `json:"teamId"`
	}

	featuredGames struct {
		GameList              []featuredGame `json:"gameList"`
		ClientRefreshInterval int64          `json:"clientRefreshInterval"`
	}
)

func (c *client) GetCurrentGameBySummonerID(ctx context.Context, region, ID string) (*currentGame, error) {
	var currentGame currentGame
	if err := c.doRequest(ctx, region, "/lol/spectator/v4/active-games/by-summoner/" + ID, &currentGame); err != nil {
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
