package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()
	Publish(`{"transaction":"$1000.00"}`, "test", producer, []byte("transaction"), deliveryChan)

	go DeliveryReport(deliveryChan)

	fmt.Println("Enviando...")
	producer.Flush(5000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "fc2-kafka-golang-kafka-1:9092",
		"delivery.timeout.ms": "0",
		"acks":                "all",
		"enable.idempotence":  "true",
	}

	producer, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Fatal(err.Error())
	}

	return producer
}

func Publish(msg, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value: []byte(msg),
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key: key,
	}

	if err := producer.Produce(message, nil); err != nil {
		return err
	}

	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Error at send message")
			} else {
				fmt.Println("Message sended: ", ev.TopicPartition)
				// save in database
			}
		}
	}
}
