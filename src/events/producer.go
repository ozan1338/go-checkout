package events

import (
	"encoding/json"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	BOOTSRAP_SERVER = "BOOTSTRAP_SERVERS"
	SERCURITY_PROTOCOL = "SECURITY_PROTOCOL"
	SASL_USERNAME = "SASL_USERNAME"
	SASL_PASSWORD = "SASL_PASSWORD"
	SASL_MECHANISM = "SASL_MECHANISM"
)

var (
	bootstrap_server = os.Getenv(BOOTSRAP_SERVER)
	security_protocol = os.Getenv(SERCURITY_PROTOCOL)
	sasl_username = os.Getenv(SASL_USERNAME)
	sasl_password = os.Getenv(SASL_PASSWORD)
	sasl_mechanism = os.Getenv(SASL_MECHANISM)
)

var Producer *kafka.Producer

func SetupProducer() {
	var err error
	Producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrap_server,
		"security.protocol": security_protocol,
		"sasl.username": sasl_username,
		"sasl.password": sasl_password,
		"sasl.mechanism": sasl_mechanism,
	})
	if err != nil {
		panic(err)
	}

	// defer Producer.Close()
}


func Produce(topic string,message interface{}, key string) {

	value, _ := json.Marshal(message)

	Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key: []byte(key),
		Value:          value,
	}, nil)

	Producer.Flush(15000)
}