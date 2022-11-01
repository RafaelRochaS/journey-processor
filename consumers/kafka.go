package consumers

import (
	"fmt"
	"log"
	"os"

	"github.com/RafaelRochaS/journey-processor/utils"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func HandleKafkaEvents() {
	consumer, err := setUpConsumer()

	if err != nil {
		log.Panic("Failed to set up consumer: ", err.Error())
	}

	consumeEvents(consumer)
}

func setUpConsumer() (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": utils.KAFKA_HOSTS,
		"group.id":          utils.KAFKA_GROUP_ID,
		"auto.offset.reset": "smallest"})

	return consumer, err
}

func consumeEvents(consumer *kafka.Consumer) {
	err := consumer.SubscribeTopics([]string{utils.KAFKA_TOPIC}, nil)

	if err != nil {
		log.Panic("Failed to subscribe to kafka topic: ", err.Error())
	}

	for run := true; run; {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			log.Println("Received kafka message: ", ev)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}

	consumer.Close()
}
