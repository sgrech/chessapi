package chessapi

import (
	"time"
)

// Player is details
type PlayerDetails struct {
	Avatar         string `json:"avatar"`
	PlayerID       int    `json:"player_id"`
	ID             string `json:"@id"`
	Url            string `json:"url"`
	Name           string `json:"name"`
	Username       string `json:"username"`
	Followers      int    `json:"followers"`
	Country        string `json:"country"`
	Location       string `json:"location"`
	LastOnline     int    `json:"last_online"`
	Joined         int    `json:"joined"`
	LastOnlineTime time.Time
	JoinedTime     time.Time
	Status         string `json:"status"`
	IsStreamer     bool   `json:"is_streamer"`
}
