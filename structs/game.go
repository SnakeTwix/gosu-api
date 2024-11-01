package structs

import "time"

type MatchGame struct {
	Id int `json:"id"`

	Beatmap Beatmap `json:"beatmap"`

	BeatmapId   int       `json:"beatmap_id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Mode        Ruleset   `json:"mode"`
	ModeInt     int       `json:"mode_int"`
	Mods        []string  `json:"mods"`
	Scores      []Score   `json:"scores"`
	ScoringType string    `json:"scoring_type"`
	TeamType    string    `json:"team_type"`
}
