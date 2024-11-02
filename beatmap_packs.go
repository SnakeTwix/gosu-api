package api

import (
	"encoding/json"
	"github.com/SnakeTwix/gosu-api/structs"
)

type GetBeatmapPacksResponse struct {
	BeatmapPacks []structs.BeatmapPack `json:"beatmap_packs"`
	Cursor       struct {
		PackId int `json:"pack_id"`
	} `json:"cursor"`

	// If you need the next page of packs, you need to supply this value to the GetBeatmapPacksQuery
	CursorString string `json:"cursor_string"`
}

type GetBeatmapPacksQuery struct {
	Type         structs.BeatmapPackType `json:"type"`
	CursorString string                  `json:"cursor_string"`
}

func (c *Client) GetBeatmapPacks(query GetBeatmapPacksQuery) (*GetBeatmapPacksResponse, error) {
	request, err := c.getRequestV2("GET", "/beatmaps/packs", nil)
	if err != nil {
		return nil, err
	}

	q := request.URL.Query()
	if query.Type != "" {
		q.Set("type", string(query.Type))
	}

	if query.CursorString != "" {
		q.Set("cursor_string", query.CursorString)
	}
	request.URL.RawQuery = q.Encode()

	response, err := c.Send(request)
	if err != nil {
		return nil, err
	}

	var beatmapPacks GetBeatmapPacksResponse
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&beatmapPacks)
	if err != nil {
		return nil, err
	}

	return &beatmapPacks, nil
}
