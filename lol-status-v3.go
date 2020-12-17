package leago

import "context"

type (
	shardStatus struct {
		Locales   []string  `json:"locales"`
		Hostname  string    `json:"hostname"`
		Name      string    `json:"name"`
		Services  []service `json:"services"`
		Slug      string    `json:"slug"`
		RegionTag string    `json:"region_tag"`
	}

	service struct {
		Name      string     `json:"name"`
		Slug      string     `json:"slug"`
		Status    string     `json:"status"`
		Incidents []incident `json:"incidents"`
	}

	incident struct {
		ID        int64     `json:"id"`
		Active    bool      `json:"active"`
		CreatedAt string    `json:"created_at"`
		Updates   []message `json:"updates"`
	}

	message struct {
		ID           string        `json:"id"`
		Author       string        `json:"author"`
		Heading      string        `json:"heading"`
		Content      string        `json:"content"`
		Severity     string        `json:"severity"`
		CreatedAt    string        `json:"created_at"`
		UpdatedAt    string        `json:"updated_at"`
		Translations []translation `json:"translations"`
	}

	translation struct {
		UpdatedAt string `json:"updated_at"`
		Locale    string `json:"locale"`
		Content   string `json:"content"`
	}
)

func (c *client) GetStatus(ctx context.Context, region string) (*shardStatus, error) {
	var status shardStatus
	if err := c.doRequest(ctx, region, "/lol/status/v3/shard-data", &status); err != nil {
		return nil, err
	}
	return &status, nil
}
