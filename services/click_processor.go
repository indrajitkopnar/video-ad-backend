package services

import (
	"time"
	"video-ad-backend/database"
	"video-ad-backend/models"
	"video-ad-backend/utils"
)

var ClickChannel = make(chan models.ClickRequest, 1000)
var BackupQueue = make(chan models.ClickRequest, 1000) // for failed DB inserts

func StartBackgroundClickProcessor() {
	go func() {
		for click := range ClickChannel {
			success := saveWithRetries(click)
			if !success {
				utils.Log.WithFields(map[string]interface{}{
					"ad_id": click.AdID,
					"ip":    click.IP,
				}).Warn("⚠️ DB write failed after retries, pushing to backup queue")
				BackupQueue <- click
			}
		}
	}()
}

// Save click with retries
func saveWithRetries(click models.ClickRequest) bool {
	var err error
	maxRetries := 3
	delay := time.Second

	for i := 0; i < maxRetries; i++ {
		err = saveClickToDB(click)
		if err == nil {
			return true
		}
		utils.Log.WithError(err).Warnf("Retry %d failed for DB insert", i+1)
		time.Sleep(delay)
		delay *= 2 // exponential backoff
	}

	return false
}

func saveClickToDB(click models.ClickRequest) error {
	query := "INSERT INTO clicks (ad_id, ip, timestamp, playback_time) VALUES ($1, $2, $3, $4)"
	_, err := database.DB.Exec(query, click.AdID, click.IP, click.Timestamp, click.PlaybackTime)
	return err
}
