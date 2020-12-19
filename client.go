package leago

import (
	"context"
	"fmt"
	json "github.com/json-iterator/go"
	"github.com/throttled/throttled/v2"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	startingAddress = "https://%s.api.riotgames.com%s"
	tokenKey    	= "X-Riot-Token"
	limitKey        = "API"
	limitQuantity   = 1
)

type client struct {
	// A client that implements the functionality of
	// interacting with League of Legends game data.

	// Riot API access token, which can only be obtained from their portal.
	// To learn more: https://developer.riotgames.com/
	AccessToken string

	// An HTTP client through which requests can be made.
	HTTPClient *http.Client

	// A field that provides a limit on the number of requests
	// in a certain period of time, which avoids the 429 error.
	Limiter *throttled.GCRARateLimiter
}

// Returns a client that allows you to interact with the Riot API.
func NewClient(accessToken string, httpClient *http.Client, limiter *throttled.GCRARateLimiter) *client {
	return &client{
		AccessToken: accessToken,
		HTTPClient:  httpClient,
		Limiter:     limiter,
	}
}

// Makes a request to the API, the result of which is the filling of the received structure.
func (c *client) doRequest(ctx context.Context, region, api string, structure interface{}) error {
	request, err := c.makeRequest(ctx, region, api)
	if err != nil {
		return err
	}
	if err := c.applyLimit(); err != nil {
		return err
	}
	
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		err, found := responseErrors[response.StatusCode]
		if !found {
			err = somethingWentWrong
		}
		return err
	}
	
	readyBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if err := response.Body.Close(); err != nil {
		return err
	}
	return json.Unmarshal(readyBody, structure)
}

// Creates a request to the API, setting an access token to successfully complete it.
func (c *client) makeRequest(ctx context.Context, region, api string) (*http.Request, error) {
	if !availabilityCheck(region) {
		return nil, regionNotSupported
	}
	request, err := http.NewRequestWithContext(
		ctx, "GET", fmt.Sprintf(startingAddress, region, api), nil,
	)
	if err != nil {
		return nil, err
	}
	request.Header.Set(tokenKey, c.AccessToken)
	return request, nil
}

// Applies the limitation, but if the restrictions when adding the limiter
// field do not match the Riot API restrictions, this method is useless.
func (c *client) applyLimit() error {
	for {
		limited, ctx, err := c.Limiter.RateLimit(limitKey, limitQuantity)
		if err != nil {
			return err
		}
		if !limited {
			return nil
		}
		time.Sleep(ctx.RetryAfter)
	}
}
