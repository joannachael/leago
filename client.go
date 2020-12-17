package leago

import (
	"context"
	json "github.com/json-iterator/go"
	"github.com/throttled/throttled/v2"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type client struct {
	// A client that implements the functionality of
	// interacting with League of Legends game data.

	// Riot API access token, which can only be obtained from their portal.
	// To learn more: https://developer.riotgames.com/
	Token string

	// An HTTP client through which requests can be made.
	HTTPClient *http.Client

	// A field that provides a limit on the number of requests
	// in a certain period of time, which avoids the 429 error.
	Limiter *throttled.GCRARateLimiter
}

// Returns a client that allows you to interact with the Riot API.
func NewClient(token string, httpClient *http.Client, limiter *throttled.GCRARateLimiter) (*client, error) {
	if token == "" {
		return nil, tokenNotReceived
	}

	cli := client{
		Token:      token,
		HTTPClient: httpClient,
		Limiter:    limiter,
	}

	return &cli, nil
}

// Makes a request to the API, the result of which is the filling of the received structure.
func (c *client) doRequest(ctx context.Context, region, address string, structure interface{}) error {
	request, err := c.makeRequest(ctx, region, address)

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

	return c.unmarshalBody(response.Body, structure)
}

// Creates a request to the API, setting an access token to successfully complete it.
func (c *client) makeRequest(ctx context.Context, region, address string) (*http.Request, error) {
	if !availabilityCheck(region) {
		return nil, regionNotSupported
	}

	requestPath := "https://" + region + ".api.riotgames.com" + address

	request, err := http.NewRequestWithContext(ctx, "GET", requestPath, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Set("X-Riot-Token", c.Token)

	return request, nil
}

// Dumps the response body into the received structure.
func (c *client) unmarshalBody(body io.ReadCloser, structure interface{}) error {
	bodyToUnmarshal, err := ioutil.ReadAll(body)

	if err != nil {
		return err
	}

	if err := body.Close(); err != nil {
		return err
	}

	return json.Unmarshal(bodyToUnmarshal, structure)
}

// Applies the limitation, but if the restrictions when adding the limiter
// field do not match the Riot API restrictions, this method is useless.
func (c *client) applyLimit() error {
	for {
		limited, ctx, err := c.Limiter.RateLimit("API", 1)

		if err != nil {
			return err
		}

		if !limited {
			return nil
		}

		time.Sleep(ctx.RetryAfter)
	}
}
