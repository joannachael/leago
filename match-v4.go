package leago

import (
	"context"
	"strconv"
)

type (
	match struct {
		GameId                int64                 `json:"gameId"`
		ParticipantIdentities []participantIdentity `json:"participantIdentities"`
		QueueId               int                   `json:"queueId"`
		GameType              string                `json:"gameType"`
		GameDuration          int64                 `json:"gameDuration"`
		Teams                 []teamStats           `json:"teams"`
		PlatformId            string                `json:"platformId"`
		GameCreation          int64                 `json:"gameCreation"`
		SeasonId              int                   `json:"seasonId"`
		GameVersion           string                `json:"gameVersion"`
		MapId                 int                   `json:"mapId"`
		GameMode              string                `json:"gameMode"`
		Participants          []participant         `json:"participants"`
	}

	participantIdentity struct {
		ParticipantId int    `json:"participantId"`
		Player        player `json:"player"`
	}

	player struct {
		Icon              int    `json:"profileIcon"`
		AccountId         string `json:"accountId"`
		MatchHistoryURI   string `json:"matchHistoryUri"`
		CurrentAccountId  string `json:"currentAccountId"`
		CurrentPlatformId string `json:"currentPlatformId"`
		SummonerName      string `json:"summonerName"`
		SummonerId        string `json:"summonerId"`
		PlatformId        string `json:"platformId"`
	}

	teamStats struct {
		TowerKills           int    `json:"towerKills"`
		RiftHeraldKills      int    `json:"riftHeraldKills"`
		FirstBlood           bool   `json:"firstBlood"`
		InhibitorKills       int    `json:"inhibitorKills"`
		Bans                 []ban  `json:"bans"`
		FirstBaron           bool   `json:"firstBaron"`
		FirstDragon          bool   `json:"firstDragon"`
		DominionVictoryScore int    `json:"dominionVictoryScore"`
		DragonKills          int    `json:"dragonKills"`
		BaronKills           int    `json:"baronKills"`
		FirstInhibitor       bool   `json:"firstInhibitor"`
		FirstTower           bool   `json:"firstTower"`
		VilemawKills         int    `json:"vilemawKills"`
		FirstRiftHerald      bool   `json:"firstRiftHerald"`
		TeamId               int    `json:"teamId"`
		Win                  string `json:"win"`
	}

	ban struct {
		ChampionId int `json:"championId"`
		PickTurn   int `json:"pickTurn"`
	}

	participant struct {
		Id                        int                 `json:"participantId"`
		ChampionId                int                 `json:"championId"`
		Runes                     []participantRune   `json:"runes"`
		Stats                     participantStats    `json:"stats"`
		TeamId                    int                 `json:"teamId"`
		Timeline                  participantTimeline `json:"timeline"`
		Spell1Id                  int                 `json:"spell1Id"`
		Spell2Id                  int                 `json:"spell2Id"`
		HighestAchievedSeasonTier string              `json:"highestAchievedSeasonTier"`
		Masteries                 []mastery           `json:"masteries"`
	}

	participantRune struct {
		Id   int `json:"runeId"`
		Rank int `json:"rank"`
	}

	participantStats struct {
		Item0                           int   `json:"item0"`
		Item1                           int   `json:"item1"`
		Item2                           int   `json:"item2"`
		Item3                           int   `json:"item3"`
		Item4                           int   `json:"item4"`
		Item5                           int   `json:"item5"`
		Item6                           int   `json:"item6"`
		TotalUnitsHealed                int   `json:"totalUnitsHealed"`
		LargestMultiKill                int   `json:"largestMultiKill"`
		GoldEarned                      int   `json:"goldEarned"`
		FirstInhibitorKill              bool  `json:"firstInhibitorKill"`
		PhysicalDamageTaken             int64 `json:"physicalDamageTaken"`
		NodeNeutralizeAssist            int   `json:"nodeNeutralizeAssist"`
		TotalPlayerScore                int   `json:"totalPlayerScore"`
		ChampLevel                      int   `json:"champLevel"`
		DamageDealtToObjectives         int64 `json:"damageDealtToObjectives"`
		TotalDamageTaken                int64 `json:"totalDamageTaken"`
		NeutralMinionsKilled            int   `json:"neutralMinionsKilled"`
		Deaths                          int   `json:"deaths"`
		TripleKills                     int   `json:"tripleKills"`
		MagicDamageDealtToChampions     int64 `json:"magicDamageDealtToChampions"`
		WardsKilled                     int   `json:"wardsKilled"`
		PentaKills                      int   `json:"pentaKills"`
		DamageSelfMitigated             int64 `json:"damageSelfMitigated"`
		LargestCriticalStrike           int   `json:"largestCriticalStrike"`
		NodeNeutralize                  int   `json:"nodeNeutralize"`
		TotalTimeCrowdControlDealt      int   `json:"totalTimeCrowdControlDealt"`
		FirstTowerKill                  bool  `json:"firstTowerKill"`
		MagicDamageDealt                int64 `json:"magicDamageDealt"`
		TotalScoreRank                  int   `json:"totalScoreRank"`
		NodeCapture                     int   `json:"nodeCapture"`
		WardsPlaced                     int   `json:"wardsPlaced"`
		TotalDamageDealt                int64 `json:"totalDamageDealt"`
		TimeCCingOthers                 int64 `json:"timeCCingOthers"`
		MagicalDamageTaken              int64 `json:"magicalDamageTaken"`
		LargestKillingSpree             int   `json:"largestKillingSpree"`
		TotalDamageDealtToChampions     int64 `json:"totalDamageDealtToChampions"`
		PhysicalDamageDealtToChampions  int64 `json:"physicalDamageDealtToChampions"`
		NeutralMinionsKilledTeamJungle  int   `json:"neutralMinionsKilledTeamJungle"`
		TotalMinionsKilled              int   `json:"totalMinionsKilled"`
		FirstInhibitorAssist            bool  `json:"firstInhibitorAssist"`
		VisionWardsBoughtInGame         int   `json:"visionWardsBoughtInGame"`
		ObjectivePlayerScore            int   `json:"objectivePlayerScore"`
		Kills                           int   `json:"kills"`
		FirstTowerAssist                bool  `json:"firstTowerAssist"`
		CombatPlayerScore               int   `json:"combatPlayerScore"`
		InhibitorKills                  int   `json:"inhibitorKills"`
		TurretKills                     int   `json:"turretKills"`
		ParticipantId                   int   `json:"participantId"`
		TrueDamageTaken                 int64 `json:"trueDamageTaken"`
		FirstBloodAssist                bool  `json:"firstBloodAssist"`
		NodeCaptureAssist               int   `json:"nodeCaptureAssist"`
		Assists                         int   `json:"assists"`
		TeamObjective                   int   `json:"teamObjective"`
		AltarsNeutralized               int   `json:"altarsNeutralized"`
		GoldSpent                       int   `json:"goldSpent"`
		DamageDealtToTurrets            int64 `json:"damageDealtToTurrets"`
		AltarsCaptured                  int64 `json:"altarsCaptured"`
		Win                             bool  `json:"win"`
		TotalHeal                       int64 `json:"totalHeal"`
		UnrealKills                     int   `json:"unrealKills"`
		VisionScore                     int64 `json:"visionScore"`
		PhysicalDamageDealt             int64 `json:"physicalDamageDealt"`
		FirstBloodKill                  bool  `json:"firstBloodKill"`
		LongestTimeSpentLiving          int   `json:"longestTimeSpentLiving"`
		KillingSprees                   int   `json:"killingSprees"`
		SightWardsBoughtInGame          int   `json:"sightWardsBoughtInGame"`
		TrueDamageDealtToChampions      int64 `json:"trueDamageDealtToChampions"`
		NeutralMinionsKilledEnemyJungle int   `json:"neutralMinionsKilledEnemyJungle"`
		DoubleKills                     int   `json:"doubleKills"`
		TrueDamageDealt                 int64 `json:"trueDamageDealt"`
		QuadraKills                     int   `json:"quadraKills"`
		PlayerScore0                    int   `json:"playerScore0"`
		PlayerScore1                    int   `json:"playerScore1"`
		PlayerScore2                    int   `json:"playerScore2"`
		PlayerScore3                    int   `json:"playerScore3"`
		PlayerScore4                    int   `json:"playerScore4"`
		PlayerScore5                    int   `json:"playerScore5"`
		PlayerScore6                    int   `json:"playerScore6"`
		PlayerScore7                    int   `json:"playerScore7"`
		PlayerScore8                    int   `json:"playerScore8"`
		PlayerScore9                    int   `json:"playerScore9"`
		Perk0                           int   `json:"perk0"`
		Perk0Var1                       int   `json:"perk0Var1"`
		Perk0Var2                       int   `json:"perk0Var2"`
		Perk0Var3                       int   `json:"perk0Var3"`
		Perk1                           int   `json:"perk1"`
		Perk1Var1                       int   `json:"perk1Var1"`
		Perk1Var2                       int   `json:"perk1Var2"`
		Perk1Var3                       int   `json:"perk1Var3"`
		Perk2                           int   `json:"perk2"`
		Perk2Var1                       int   `json:"perk2Var1"`
		Perk2Var2                       int   `json:"perk2Var2"`
		Perk2Var3                       int   `json:"perk2Var3"`
		Perk3                           int   `json:"perk3"`
		Perk3Var1                       int   `json:"perk3Var1"`
		Perk3Var2                       int   `json:"perk3Var2"`
		Perk3Var3                       int   `json:"perk3Var3"`
		Perk4                           int   `json:"perk4"`
		Perk4Var1                       int   `json:"perk4Var1"`
		Perk4Var2                       int   `json:"perk4Var2"`
		Perk4Var3                       int   `json:"perk4Var3"`
		Perk5                           int   `json:"perk5"`
		Perk5Var1                       int   `json:"perk5Var1"`
		Perk5Var2                       int   `json:"perk5Var2"`
		Perk5Var3                       int   `json:"perk5Var3"`
		PerkPrimaryStyle                int   `json:"perkPrimaryStyle"`
		PerkSubStyle                    int   `json:"perkSubStyle"`
		StatPerk0                       int   `json:"statPerk0"`
		StatPerk1                       int   `json:"statPerk1"`
		StatPerk2                       int   `json:"statPerk2"`
	}

	participantTimeline struct {
		ParticipantId               int                `json:"participantId"`
		CsDiffPerMinDeltas          map[string]float32 `json:"cSDiffPerMinDeltas"`
		DamageTakenPerMinDeltas     map[string]float32 `json:"damageTakenPerMinDeltas"`
		Role                        string             `json:"role"`
		DamageTakenDiffPerMinDeltas map[string]float32 `json:"damageTakenDiffPerMinDeltas"`
		XpPerMinDeltas              map[string]float32 `json:"xPPerMinDeltas"`
		XpDiffPerMinDeltas          map[string]float32 `json:"xPDiffPerMinDeltas"`
		Lane                        string             `json:"lane"`
		CreepsPerMinDeltas          map[string]float32 `json:"creepsPerMinDeltas"`
		GoldPerMinDeltas            map[string]float32 `json:"goldPerMinDeltas"`
	}

	mastery struct {
		Rank int `json:"rank"`
		Id   int `json:"masteryId"`
	}

	matches struct {
		StartIndex int              `json:"startIndex"`
		TotalGames int              `json:"totalGames"`
		EndIndex   int              `json:"endIndex"`
		Matches    []matchReference `json:"matches"`
	}

	matchReference struct {
		GameId     int64  `json:"gameId"`
		Role       string `json:"role"`
		Season     int    `json:"season"`
		PlatformId string `json:"platformId"`
		Champion   int    `json:"champion"`
		Queue      int    `json:"queue"`
		Lane       string `json:"lane"`
		TimeStamp  int64  `json:"timestamp"`
	}

	matchTimeline struct {
		Frames        []frame `json:"frames"`
		FrameInterval int64   `json:"frameInterval"`
	}

	frame struct {
		ParticipantFrames map[string]participantFrame `json:"participantFrames"`
		Events            []event                     `json:"events"`
		Timestamp         int64                       `json:"timestamp"`
	}

	participantFrame struct {
		ParticipantId       int      `json:"participantId"`
		MinionsKilled       int      `json:"minionsKilled"`
		TeamScore           int      `json:"teamScore"`
		DominionScore       int      `json:"dominionScore"`
		TotalGold           int      `json:"totalGold"`
		Level               int      `json:"level"`
		Xp                  int      `json:"xp"`
		CurrentGold         int      `json:"currentGold"`
		Position            position `json:"position"`
		JungleMinionsKilled int      `json:"jungleMinionsKilled"`
	}

	event struct {
		LaneType                string   `json:"laneType"`
		SkillSlot               int      `json:"skillSlot"`
		AscendedType            string   `json:"ascendedType"`
		CreatorId               int      `json:"creatorId"`
		AfterId                 int      `json:"afterId"`
		EventType               string   `json:"eventType"`
		Type                    string   `json:"type"`
		LevelUpType             string   `json:"levelUpType"`
		WardType                string   `json:"wardType"`
		ParticipantId           int      `json:"participantId"`
		TowerType               string   `json:"towerType"`
		ItemId                  int      `json:"itemId"`
		BeforeId                int      `json:"beforeId"`
		PointCaptured           string   `json:"pointCaptured"`
		MonsterType             string   `json:"monsterType"`
		MonsterSubType          string   `json:"monsterSubType"`
		TeamId                  int      `json:"teamId"`
		Position                position `json:"position"`
		KillerId                int      `json:"killerId"`
		Timestamp               int64    `json:"timestamp"`
		AssistingParticipantIds []int    `json:"assistingParticipantIds"`
		BuildingType            string   `json:"buildingType"`
		VictimId                int      `json:"victimId"`
	}

	position struct {
		X int `json:"x"`
		Y int `json:"y"`
	}
)

func (c *client) GetMatchById(ctx context.Context, region string, matchId int64) (*match, error) {
	var match match
	if err := c.doRequest(ctx, region, "/lol/match/v4/matches/" + strconv.FormatInt(matchId, 10), &match); err != nil {
		return nil, err
	}
	return &match, nil
}

func (c *client) GetMatchByIdAndTournamentCode(ctx context.Context, region, tournamentCode string, matchId int64) (*match, error) {
	var match match
	address := "/lol/match/v4/matches/" + strconv.FormatInt(matchId, 10) + "/by-tournament-code/" + tournamentCode
	if err := c.doRequest(ctx, region, address, &match); err != nil {
		return nil, err
	}
	return &match, nil
}

func (c *client) GetMatchesBySummonerAccountId(ctx context.Context, region, accountId string) (*matches, error) {
	var matches matches
	if err := c.doRequest(ctx, region, "/lol/match/v4/matchlists/by-account/" + accountId, &matches); err != nil {
		return nil, err
	}
	return &matches, nil
}

func (c *client) GetMatchIdsByTournamentCode(ctx context.Context, region, tournamentCode string) ([]int64, error) {
	var matchIds []int64
	address := "/lol/match/v4/matches/by-tournament-code/" + tournamentCode + "/ids"
	if err := c.doRequest(ctx, region, address, &matchIds); err != nil {
		return nil, err
	}
	return matchIds, nil
}

func (c *client) GetTimelinesByMatchId(ctx context.Context, region string, matchId int64) (*matchTimeline, error) {
	var timelines matchTimeline
	address := "/lol/match/v4/timelines/by-match/" + strconv.FormatInt(matchId, 10)
	if err := c.doRequest(ctx, region, address, &timelines); err != nil {
		return nil, err
	}
	return &timelines, nil
}
