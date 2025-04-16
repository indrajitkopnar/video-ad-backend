//video-ad-backend/controllers/ad_controller.go

package controllers

import (
	"fmt"
	"net/http"
	"time"
	"video-ad-backend/database"
	"video-ad-backend/models"
	"video-ad-backend/services"
	"video-ad-backend/utils"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/gin-gonic/gin"
)

type Ad struct {
	ID        int    `json:"id"`
	ImageURL  string `json:"image_url"`
	TargetURL string `json:"target_url"`
}

func GetAds(c *gin.Context) {
	utils.HttpRequestsTotal.WithLabelValues("/ads", "GET").Inc()
	timer := prometheus.NewTimer(utils.HttpRequestDuration.WithLabelValues("/ads"))
	defer timer.ObserveDuration()

	rows, err := database.DB.Query("SELECT id, image_url, target_url FROM ads")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch ads"})
		return
	}
	defer rows.Close()

	var ads []Ad
	for rows.Next() {
		var ad Ad
		if err := rows.Scan(&ad.ID, &ad.ImageURL, &ad.TargetURL); err != nil {
			continue
		}
		ads = append(ads, ad)
	}

	c.JSON(http.StatusOK, ads)
}

func PostClick(c *gin.Context) {
	utils.HttpRequestsTotal.WithLabelValues("/ads/click", "POST").Inc()
	timer := prometheus.NewTimer(utils.HttpRequestDuration.WithLabelValues("/ads/click"))
	defer timer.ObserveDuration()

	var click models.ClickRequest
	if err := c.ShouldBindJSON(&click); err != nil {
		utils.Log.WithField("error", err.Error()).Warn("❌ Invalid click payload")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Parse timestamp from request
	parsedTime, err := time.Parse(time.RFC3339, click.Timestamp)
	if err != nil {
		utils.Log.WithError(err).Error("❌ Timestamp parse error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid timestamp format"})
		return
	}

	ip := c.ClientIP()
	click.IP = ip
	truncated := parsedTime.Truncate(time.Minute).Unix()
	dedupKey := fmt.Sprintf("dedup:%d:%s:%d", click.AdID, ip, truncated)

	// Deduplication in Redis
	set, err := database.RedisClient.SetNX(database.RedisCtx, dedupKey, 1, 10*time.Minute).Result()
	if err != nil {
		utils.Log.WithError(err).Error("❌ Redis SETNX failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}
	if !set {
		utils.Log.WithFields(map[string]interface{}{
			"ad_id":    click.AdID,
			"ip":       ip,
			"dedupKey": dedupKey,
		}).Info("⚠️ Duplicate click ignored")
		c.JSON(http.StatusOK, gin.H{"message": "Duplicate click ignored"})
		return
	}

	// Send to processor channel
	go func() {
		services.ClickChannel <- click
	}()

	utils.Log.WithFields(map[string]interface{}{
		"ad_id":    click.AdID,
		"ip":       ip,
		"datetime": parsedTime.Format(time.RFC3339),
	}).Info("✅ Click accepted")

	c.JSON(http.StatusOK, gin.H{"message": "Click received"})
}

func GetAdAnalytics(c *gin.Context) {
	adIDs := []int{1, 2, 3, 4, 5}

	type Analytics struct {
		AdID        int     `json:"ad_id"`
		ClickCount  int     `json:"click_count"`
		Impressions int     `json:"impressions"`
		CTR         float64 `json:"ctr"`
	}

	var results []Analytics

	for _, id := range adIDs {
		clicksStr, err := database.RedisClient.Get(database.RedisCtx, fmt.Sprintf("ad_clicks:%d", id)).Result()
		if err != nil {
			clicksStr = "0"
		}
		impressionsStr, err := database.RedisClient.Get(database.RedisCtx, fmt.Sprintf("ad_impressions:%d", id)).Result()
		if err != nil {
			impressionsStr = "0"
		}

		var clicks, impressions int
		fmt.Sscanf(clicksStr, "%d", &clicks)
		fmt.Sscanf(impressionsStr, "%d", &impressions)

		ctr := 0.0
		if impressions > 0 {
			ctr = (float64(clicks) / float64(impressions)) * 100
		}

		results = append(results, Analytics{
			AdID:        id,
			ClickCount:  clicks,
			Impressions: impressions,
			CTR:         ctr,
		})
	}

	c.JSON(http.StatusOK, results)
}
