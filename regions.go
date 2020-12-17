package leago

const (
	AMERICAS = "americas"
	ASIA     = "asia"
	EUROPE   = "europe"
	BR1      = "br1"
	EUN1     = "eun1"
	EUW1     = "euw1"
	JP1      = "jp1"
	KR       = "kr"
	LA1      = "la1"
	LA2      = "la2"
	NA1      = "na1"
	OC1      = "oc1"
	TR1      = "tr1"
	RU       = "ru"
)

var regions = []string{
	AMERICAS, ASIA, EUROPE, BR1, EUN1, EUW1, JP1, KR, LA1, LA2, NA1, OC1, TR1, RU,
}

func availabilityCheck(region string) bool {
	for _, r := range regions {
		if region == r {
			return true
		}
	}
	return false
}
