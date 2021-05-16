package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"github.com/matheuslss/code-delivery/simulator/infra/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChannel := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChannel)

	go consumer.Consume()

	for msg := range msgChannel {
		fmt.Println(string(msg.Value))
	}
}
