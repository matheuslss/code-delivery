package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	simRoute "github.com/matheuslss/code-delivery/simulator/application/route"
	"github.com/matheuslss/code-delivery/simulator/infra/kafka"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := simRoute.NewRoute()
	json.Unmarshal(msg.Value, &route)

	route.LoadPositions()

	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}

	for _, position := range positions {
		kafka.Publish(position, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
