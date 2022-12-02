package main

import (
	"fmt"
	"go-checkout/src/database"
	"go-checkout/src/events"
	"go-checkout/src/models"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	BOOTSRAP_SERVER = "BOOTSTRAP_SERVERS"
	SERCURITY_PROTOCOL = "SECURITY_PROTOCOL"
	SASL_USERNAME = "SASL_USERNAME"
	SASL_PASSWORD = "SASL_PASSWORD"
	SASL_MECHANISM = "SASL_MECHANISM"
	KAFKA_TOPIC = "KAFKA_TOPIC"
)

var (
	bootstrap_server = os.Getenv(BOOTSRAP_SERVER)
	security_protocol = os.Getenv(SERCURITY_PROTOCOL)
	sasl_username = os.Getenv(SASL_USERNAME)
	sasl_password = os.Getenv(SASL_PASSWORD)
	sasl_mechanism = os.Getenv(SASL_MECHANISM)
	kafka_topic = os.Getenv(KAFKA_TOPIC)
)

func main() {
	database.Connect()
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrap_server,
		"security.protocol": security_protocol,
		"sasl.username": sasl_username,
		"sasl.password": sasl_password,
		"sasl.mechanism": sasl_mechanism,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	fmt.Println("START CONSUMING")

	if err != nil {
		panic(err)
	}

	consumer.SubscribeTopics([]string{kafka_topic}, nil)

	fmt.Println(kafka_topic)

	defer consumer.Close()

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			database.DB.Create(&models.KafkaError{
				Key: msg.Key,
				Value: msg.Value,
				Error: err,
			})
			return
		}

		fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		if err := events.Listen(msg); err != nil {
			database.DB.Create(&models.KafkaError{
				Key: msg.Key,
				Value: msg.Value,
				Error: err,
			})
		}
	}
}