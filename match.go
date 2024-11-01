package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SnakeTwix/gosu-api/structs"
	"strconv"
)

type MatchData struct {
	Match  structs.Match        `json:"match"`
	Events []structs.MatchEvent `json:"events"`

	Users []structs.User `json:"users"`

	FirstEventId  int `json:"first_event_id"`
	LatestEventId int `json:"latest_event_id"`
}

type GetMatchQuery struct {
	MatchId int
	Before  int
	After   int
	Limit   int
}

func (c *Client) GetMatch(query GetMatchQuery) (*MatchData, error) {
	if query.MatchId == 0 {
		return nil, errors.New("no id provided for match query")
	}

	request, err := c.getRequestV2("GET", fmt.Sprintf("/matches/%d", query.MatchId), nil)
	if err != nil {
		return nil, err
	}

	q := request.URL.Query()
	if query.Limit != 0 {
		q.Set("limit", strconv.Itoa(query.Limit))
	}

	if query.After != 0 {
		q.Set("after", strconv.Itoa(query.After))
	}

	if query.Before != 0 {
		q.Set("before", strconv.Itoa(query.Before))
	}

	request.URL.RawQuery = q.Encode()

	response, err := c.Send(request)
	if err != nil {
		return nil, err
	}

	var matchData MatchData
	decoder := json.NewDecoder(response.Body)
	decoder.UseNumber()
	err = decoder.Decode(&matchData)
	if err != nil {
		return nil, err
	}

	return &matchData, nil
}

func (c *Client) GetFullMatch(id int) (*MatchData, error) {
	matchQuery := GetMatchQuery{MatchId: id}

	matchData, err := c.GetMatch(matchQuery)
	if err != nil {
		return nil, err
	}

	trackedUsers := map[int]bool{}
	for userIndex := range matchData.Users {
		user := &matchData.Users[userIndex]

		trackedUsers[user.Id] = true
	}

	events := matchData.Events

	// Fetch until we get all the events and users for a match
	for events[0].Id != matchData.FirstEventId {
		matchQuery.Before = events[0].Id
		tempMatchData, err := c.GetMatch(matchQuery)
		if err != nil {
			return nil, err
		}

		events = append(tempMatchData.Events, events...)

		for userIndex := range tempMatchData.Users {
			user := &tempMatchData.Users[userIndex]
			if !trackedUsers[user.Id] {
				trackedUsers[user.Id] = true
				matchData.Users = append(matchData.Users, *user)
			}
		}
	}

	matchData.Events = events
	return matchData, nil
}
