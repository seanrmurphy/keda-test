package main

import (
	"context"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
)

const (
	MaxSleepTime          int    = 20
	ReceiveTopic          string = "sysbench-test.sbtest.sbtest1"
	ReceiveTopicPartition int    = 0
	ConsumerGroup         string = "knative-group"
	KafkaCluster          string = "stooling-cluster-kafka-bootstrap.kafka:9092"
	//RespondTopic          string = "sysbench-test.sbtest.sbtest1"
	//RespondTopicPartition int    = 0
)

type Message struct {
	s string
}

func ParseMessage(msg []byte) (m Message) {
	m.s = string(msg)
	return
}

// main entry point
func main() {
	// start kafka consumer and listen...

	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{KafkaCluster},
		GroupID:  ConsumerGroup,
		Topic:    ReceiveTopic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	r.Close()

}
