package main

import (
	"fmt"

	cyclopskafka "gitlab.com/cyclops-community/bill-notification/test"
	l "gitlab.com/cyclops-utilities/logging"
)

const (
	MaxSleepTime          int    = 20
	ReceiveTopic          string = "sysbench-test.sbtest.sbtest1"
	ReceiveTopicPartition int    = 0
	//RespondTopic          string = "sysbench-test.sbtest.sbtest1"
	//RespondTopicPartition int    = 0
	KafkaCluster string = "stooling-cluster-kafka-bootstrap.kafka:9092"
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

	kafkaReceiver := cyclopskafka.KafkaHandler{
		Broker: KafkaCluster,
		Topic:  ReceiveTopic,
	}

	kafkaReceiver.Initialize()

	counter := int64(0)
	for {
		err, key, val := kafkaReceiver.ReadMessage(counter)
		if err != nil {
			l.Error.Printf("Error reading message from kafka: %v\n", err.Error())
		}
		counter++

		fmt.Printf("message at offset %v: %v = %v\n", counter, string(key), string(val))

		receivedMessage := ParseMessage(val)

		l.Info.Printf("receivedMessage = %v\n", receivedMessage)
	}

	// never happens
	//r.Close()
}
