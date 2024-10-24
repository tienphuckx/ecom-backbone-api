package initialize

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/tienphuckx/ecom-backbone-api.git/global"
	"go.uber.org/zap"
)

var kafkaWriter *kafka.Writer

// InitKafka initializes the Kafka producer
func InitKafka() error {
	kafkaConfig := global.SysConfig.KafkaConfig

	// Create a new Kafka writer (producer)
	kafkaWriter = &kafka.Writer{
		Addr:         kafka.TCP(kafkaConfig.Brokers...),
		Topic:        kafkaConfig.Topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
	}

	// Test Kafka connection by sending a test message (optional)
	msg := kafka.Message{
		Key:   []byte("init-test-key"),
		Value: []byte("init-test-message"),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := kafkaWriter.WriteMessages(ctx, msg)
	if err != nil {
		return fmt.Errorf("failed to connect to Kafka: %w", err)
	}

	global.Logger.Info("Kafka producer successfully initialized", zap.Strings("brokers", kafkaConfig.Brokers))
	return nil
}
