package structs

import "time"

type BeatmapPack struct {
	Author          string    `json:"author"`
	Date            time.Time `json:"date"`
	Name            string    `json:"name"`
	NoDiffReduction bool      `json:"no_diff_reduction"`
	RulesetId       int       `json:"ruleset_id"`
	Tag             string    `json:"tag"`
	Url             string    `json:"url"`

	Beatmapsets        []BeatmapSet `json:"beatmapsets"`
	UserCompletionData *struct {
		BeatmapSetIds []int `json:"beatmap_set_ids"`
		Completed     bool  `json:"completed"`
	} `json:"user_completion_data"`
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
