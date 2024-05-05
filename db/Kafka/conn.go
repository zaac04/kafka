package Kafka

import (
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

var client *kafka.Conn

func Connect() {
	var err error
	client, err = kafka.Dial("tcp", os.Getenv("KAFKA_URL"))

	if err != nil {
		fmt.Println(err)
	}
}

func Close() {
	err := client.Close()
	if err != nil {
		fmt.Println(err)
	}
}
