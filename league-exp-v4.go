package leago

import (
	"context"
	"strconv"
)

func (c *client) GetEntriesByQueueTierAndDivisionEXP(
	ctx context.Context, region, queue, tier, division string, page int,
) ([]entry, error) {
	var entries []entry
	address := "/lol/league-exp/v4/entries/" + queue + "/" + tier + "/" + division + "?page=" + strconv.Itoa(page)
	if err := c.doRequest(ctx, region, address, &entries); err != nil {
		return nil, err
	}
	return entries, nil
}
