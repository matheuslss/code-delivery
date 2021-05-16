package kafka

import (
	"encoding/json"
	"os"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	simRoute "github.com/matheuslss/code-delivery/simulator/application/route"
	"github.com/matheuslss/code-delivery/simulator/infra/kafka"
	"github.com/sirupsen/logrus"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := simRoute.NewRoute()
	err := json.Unmarshal(msg.Value, &route)
	if err != nil {
		logrus.Error(err.Error())
	}

	err = route.LoadPositions()
	if err != nil {
		logrus.Error(err.Error())
	}

	positions, err := route.ExportJsonPositions()
	if err != nil {
		logrus.Error(err.Error())
	}

	for _, position := range positions {
		kafka.Publish(position, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
