//video-ad-backend/models/ad.go

package models

type AdAnalytics struct {
	AdID        int `json:"ad_id"`
	ClickCount  int `json:"click_count"`
	Impressions int `json:"impressions"`
	CTR         int `json:"ctr"` // Rounded percentage
}
