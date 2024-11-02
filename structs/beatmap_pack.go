package structs

import "time"

type BeatmapPack struct {
	Author          string    `json:"author"`
	Date            time.Time `json:"date"`
	Name            string    `json:"name"`
	NoDiffReduction bool      `json:"no_diff_reduction"`
	RulesetId       int       `json:"ruleset_id"`
	// Maybe needs typing
	Tag string `json:"tag"`

	Url string `json:"url"`
}

type BeatmapPackType string

const (
	BeatmapPackStandard   BeatmapPackType = "standard"
	BeatmapPackFeatured   BeatmapPackType = "featured"
	BeatmapPackTournament BeatmapPackType = "tournament"
	BeatmapPackLoved      BeatmapPackType = "loved"
	BeatmapPackChart      BeatmapPackType = "chart"
	BeatmapPackTheme      BeatmapPackType = "theme"
	BeatmapPackArtist     BeatmapPackType = "artist"
)
