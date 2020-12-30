package leago

import "context"

type (
	platformData struct {
		Id           string   `json:"id"`
		Name         string   `json:"name"`
		Locales      []string `json:"locales"`
		Maintenances []status `json:"maintenances"`
		Incidents    []status `json:"incidents"`
	}

	status struct {
		Id                int       `json:"id"`
		MaintenanceStatus string    `json:"maintenance_status"`
		IncidentSeverity  string    `json:"incident_severity"`
		Titles            []content `json:"titles"`
		Updates           []update  `json:"updates"`
		CreatedAt         string    `json:"created_at"`
		ArchiveAt         string    `json:"archive_at"`
		UpdatedAt         string    `json:"updated_at"`
		Platforms         []string  `json:"platforms"`
	}

	update struct {
		Id               int       `json:"id"`
		Author           string    `json:"author"`
		Publish          bool      `json:"publish"`
		PublishLocations []string  `json:"publish_locations"`
		Translations     []content `json:"translations"`
		CreatedAt        string    `json:"created_at"`
		UpdatedAt        string    `json:"updated_at"`
	}

	content struct {
		Locale  string `json:"locale"`
		Content string `json:"content"`
	}
)

func (c *client) GetStatus(ctx context.Context, region string) (*platformData, error) {
	var platformData platformData
	if err := c.doRequest(ctx, region, "/lol/status/v4/platform-data", &platformData); err != nil {
		return nil, err
	}
	return &platformData, nil
}
