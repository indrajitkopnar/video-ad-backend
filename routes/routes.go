package routes

import (
	"video-ad-backend/controllers"
	"video-ad-backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/ads", controllers.GetAds)
	router.POST("/ads/click", middleware.RateLimitMiddleware(), controllers.PostClick)
	router.GET("/ads/analytics", controllers.GetAdAnalytics)
}
