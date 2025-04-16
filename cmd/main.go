//video-ad-backend/cmd/main.go

package main

import (
	"fmt"
	"os"

	"video-ad-backend/config"
	"video-ad-backend/database"
	"video-ad-backend/routes"
	"video-ad-backend/services"

	"video-ad-backend/utils"

	"video-ad-backend/kafka"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	utils.InitLogger()
	database.InitPostgres()
	database.InitRedis()
	utils.InitMetrics()
	services.StartBackgroundClickProcessor()
	kafka.InitKafkaProducer()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	r := gin.Default()
	routes.RegisterRoutes(r)

	fmt.Println("🚀 Server running on port " + port)
	r.Run(":" + port)
}
