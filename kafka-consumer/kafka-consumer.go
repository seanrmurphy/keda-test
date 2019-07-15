package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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

type aMessage struct {
	CreatedAt string `json:"created_at"`
	ID        int    `json:"id"`
}

type pMessage struct {
	After     aMessage `json:"after"`
	Timestamp int      `json:"ts_ms"`
}

type ParsedMessage struct {
	Payload pMessage `json:"payload"`
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
		val := m.Value
		p := ParsedMessage{}
		json.Unmarshal(val, &p)
		unixTime := time.Now().Unix()
		payloadTimestamp := p.Payload.Timestamp
		createdAt := p.Payload.After.CreatedAt

		formatString := "2006-01-02T15:04:05.000Z"
		createdAtTime, _ := time.Parse(formatString, createdAt)

		fmt.Printf("Record ID:%v,  Time(current): %v, payload ts_ms: %v, created_at: %v, ID: %v\n",
			p.Payload.After.ID, unixTime, payloadTimestamp, createdAtTime.Unix())
	}

	r.Close()

}
