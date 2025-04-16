package models

type ClickRequest struct {
	AdID         int    `json:"ad_id" binding:"required"`
	IPAddress    string `json:"ip_address"`
	PlaybackTime int    `json:"video_playback_time"`
	Timestamp    string `json:"timestamp"` // Optional if using server time
	IP           string `json:"-"`         // added by server
}
