package structs

import "time"

type MatchEvent struct {
	Id        int              `json:"id"`
	Detail    MatchEventDetail `json:"detail"`
	Timestamp time.Time        `json:"timestamp"`
	UserId    *int             `json:"user_id"`
	Game      *MatchGame       `json:"game"` // Optional
}

type MatchEventDetail struct {
	Type MatchEventType `json:"type"`

	// Text seems to be only present on MatchEventOther
	Text *string `json:"text"`
}

type MatchEventType string

const (
	MatchEventHostChanged    MatchEventType = "host-changed"
	MatchEventMatchCreated   MatchEventType = "match-created"
	MatchEventMatchDisbanded MatchEventType = "match-disbanded"
	MatchEventPlayerJoined   MatchEventType = "player-joined"
	MatchEventPlayerKicked   MatchEventType = "player-kicked"
	MatchEventPlayerLeft     MatchEventType = "player-left"
	MatchEventOther          MatchEventType = "other"
)

type Match struct {
	Id        int       `json:"id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Name      string    `json:"name"`
}
