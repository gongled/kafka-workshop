package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/segmentio/kafka-go"
)

var (
	brokers = ""
	group   = ""
	topic   = ""
)

func init() {
	flag.StringVar(&brokers, "brokers", os.Getenv("KAFKA_BROKERS"), "Kafka bootstrap brokers to connect to, as a comma separated list")
	flag.StringVar(&group, "group", os.Getenv("KAFKA_CONSUMER_GROUP"), "Kafka consumer group definition")
	flag.StringVar(&topic, "topic", os.Getenv("KAFKA_TOPIC"), "Kafka topic to be consumed")
	flag.Parse()

	if len(brokers) == 0 {
		panic("no Kafka bootstrap brokers defined, please set the -brokers flag")
	}

	if len(topic) == 0 {
		panic("no topic given to be consumed, please set the -topic flag")
	}

	if len(group) == 0 {
		panic("no Kafka consumer group defined, please set the -group flag")
	}
}

func main() {
	// make a new reader that consumes from topic-A
	addrs := strings.Split(brokers, ",")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  addrs,
		GroupID:  group,
		Topic:    topic,
		MinBytes: 10e2, // 1KB
		MaxBytes: 10e6, // 10MB
	})

	log.Print("Starting consumer program...")
	log.Print(fmt.Sprintf("Brokers (%s), topic (%s), consumer group (%s)", brokers, topic, group))

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		log.Print(fmt.Sprintf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value)))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
