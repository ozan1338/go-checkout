package main

// import (
// 	"fmt"
// 	"go-admin/src/database"
// 	"go-admin/src/events"
// 	"go-admin/src/models"

// 	"github.com/confluentinc/confluent-kafka-go/kafka"
// )

// func main() {
// 	database.Connect()
// 	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
// 		"bootstrap.servers": "pkc-ew3qg.asia-southeast2.gcp.confluent.cloud:9092",
// 		"security.protocol": "SASL_SSL",
// 		"sasl.username":     "ZRITSDHTM4YORCX3",
// 		"sasl.password":     "iOJVSZ5sHVRnmunF7VvCw+lC1iADXyNZGeYuVlZZfUlvcvUn4fotwbsxRoW2WY2W",
// 		"sasl.mechanism":    "PLAIN",
// 		"group.id":          "myGroup",
// 		"auto.offset.reset": "earliest",
// 	})
// 	fmt.Println("START CONSUMING")

// 	if err != nil {
// 		panic(err)
// 	}

// 	consumer.SubscribeTopics([]string{"admin_topic"}, nil)

// 	fmt.Println("admin_topic")

// 	defer consumer.Close()

// 	for {
// 		msg, err := consumer.ReadMessage(-1)
// 		if err != nil {
// 			// The client will automatically try to recover from all errors.
// 			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
// 			database.DB.Create(&models.KafkaError{
// 				Key: msg.Key,
// 				Value: msg.Value,
// 				Error: err,
// 			})
// 			return
// 		}

// 		fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))

// 		if err := events.Listen(msg); err != nil {
// 			database.DB.Create(&models.KafkaError{
// 				Key: msg.Key,
// 				Value: msg.Value,
// 				Error: err,
// 			})
// 		}
// 	}
// }