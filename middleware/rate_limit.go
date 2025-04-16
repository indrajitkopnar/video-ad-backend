package middleware

import (
	"net/http"
	"time"
	"video-ad-backend/database"
	"video-ad-backend/utils"

	"github.com/gin-gonic/gin"
)

const rateLimit = 10
const window = time.Minute

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "rl:" + ip

		// Increment count with TTL
		count, err := database.RedisClient.Incr(database.RedisCtx, key).Result()
		if err != nil {
			utils.Log.WithError(err).Error("Redis INCR failed in rate limiter")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		if count == 1 {
			database.RedisClient.Expire(database.RedisCtx, key, window)
		}

		if count > rateLimit {
			utils.Log.WithField("ip", ip).Warn("Rate limit exceeded")
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded. Try again later."})
			return
		}

		c.Next()
	}
}
