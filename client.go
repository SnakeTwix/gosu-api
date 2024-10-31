package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SnakeTwix/gosu-api/internal/util"
	"io"
	"net/http"
	"time"
)

const BaseUrl = "https://osu.ppy.sh/api/v2"

type Client struct {
	url          string
	clientSecret string
	clientId     int
	authToken    *token
	httpClient   *http.Client
}

type token struct {
	access    string
	refresh   *string
	expiresAt time.Time
}

func NewClient(clientId int, clientSecret string) (Client, error) {
	client := Client{
		url:          BaseUrl,
		clientSecret: clientSecret,
		clientId:     clientId,
		httpClient:   http.DefaultClient,
	}

	err := client.fetchToken()
	if err != nil {
		return Client{}, err
	}

	return client, nil
}

func (c *Client) fetchToken() error {
	// TODO: get rid of body Map mappings
	content := map[string]any{
		"client_id":     c.clientId,
		"client_secret": c.clientSecret,
		"grant_type":    "client_credentials",
		"scope":         "public",
	}

	body, err := util.MapToReader(content)
	if err != nil {
		return err
	}

	response, err := c.httpClient.Post("https://osu.ppy.sh/oauth/token", "application/json", body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("token request failed")
	}

	var tokenMap map[string]any
	decoder := json.NewDecoder(response.Body)
	decoder.UseNumber()
	err = decoder.Decode(&tokenMap)
	if err != nil {
		return err
	}

	accessToken, ok := tokenMap["access_token"]
	if !ok {
		return errors.New("no access_token when requesting auth")
	}

	stringToken, ok := accessToken.(string)
	if !ok {
		return errors.New("accessToken isn't string")
	}

	expiresIn, ok := tokenMap["expires_in"]
	if !ok {
		return errors.New("no expires_in when requesting auth")
	}

	expiresJsonNumber, ok := expiresIn.(json.Number)
	if !ok {
		return errors.New("expires_in isn't number string")
	}

	expiresNumber, err := expiresJsonNumber.Int64()
	if err != nil {
		return errors.New("expires_in isn't json int")
	}

	c.authToken = &token{
		access:  stringToken,
		refresh: nil,
		expiresAt: time.Now().
			// Just in case it timeouts or something
			Add(-time.Minute * 5).
			Add(time.Second * time.Duration(expiresNumber)),
	}

	return nil
}

// Spits out a request object prepended with the v2 api url and sets the Bearer token
func (c *Client) getRequestV2(method string, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.url, url), body)
	if err != nil {
		return nil, err
	}

	if c.authToken == nil {
		return nil, errors.New("no auth token specified")
	}

	// If the token has expired
	if c.authToken.expiresAt.Before(time.Now()) {
		err = c.fetchToken()
		if err != nil {
			return nil, err
		}
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.authToken.access))

	return request, nil
}

func (c *Client) Send(request *http.Request) (*http.Response, error) {
	return c.httpClient.Do(request)
}
