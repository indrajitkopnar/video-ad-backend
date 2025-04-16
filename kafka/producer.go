package kafka

import (
	"encoding/json"
	"os"
	"time"
	"video-ad-backend/models"
	"video-ad-backend/utils"

	"github.com/segmentio/kafka-go"
)

var Writer *kafka.Writer

func InitKafkaProducer() {
	Writer = &kafka.Writer{
		Addr:     kafka.TCP(os.Getenv("KAFKA_BROKER")),
		Topic:    "video_clicks",
		Balancer: &kafka.LeastBytes{},
	}
	utils.Log.Info("✅ Kafka producer initialized")
}

func PublishClickEvent(click models.ClickRequest) {
	data, err := json.Marshal(click)
	if err != nil {
		utils.Log.WithError(err).Error("❌ Failed to serialize click event")
		return
	}

	msg := kafka.Message{
		Key:   []byte(click.IP),
		Value: data,
		Time:  time.Now(),
	}

	if err := Writer.WriteMessages(utils.Ctx, msg); err != nil {
		utils.Log.WithError(err).Error("❌ Kafka publish failed")
	} else {
		utils.Log.WithField("ad_id", click.AdID).Info("📤 Click event sent to Kafka")
	}
}
