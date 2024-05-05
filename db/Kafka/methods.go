package Kafka

import (
	"context"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

func CreateTopic(topicName string) (err error) {
	topic := kafka.TopicConfig{
		Topic:             topicName,
		NumPartitions:     1,
		ReplicationFactor: 1,
	}
	err = client.CreateTopics(topic)
	if err == kafka.TopicAlreadyExists {
		return nil
	}
	return err
}

func PublishMessage(Topic string, Message string) (err error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{os.Getenv("KAFKA_URL")},
		Topic:        Topic,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 10 * time.Millisecond,
	})

	defer writer.Close()

	message := kafka.Message{
		Value: []byte(Message),
	}
	err = writer.WriteMessages(context.Background(), message)
	if err != nil {
		return err
	}
	return
}

func GetReader(Topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{os.Getenv("KAFKA_URL")},
		Topic:     Topic,
		Partition: 0,    // Optionally specify a partition
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
}
