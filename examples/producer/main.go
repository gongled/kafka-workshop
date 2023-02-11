package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/segmentio/kafka-go"
)

var (
	brokers = ""
	topic   = ""
)

type CandidateLocation struct {
	ID  int     `json:"id"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func init() {
	flag.StringVar(&brokers, "brokers", os.Getenv("KAFKA_BROKERS"), "Kafka bootstrap brokers to connect to, as a comma separated list")
	flag.StringVar(&topic, "topic", os.Getenv("KAFKA_TOPIC"), "Kafka topic to be produced")
	flag.Parse()

	if len(brokers) == 0 {
		panic("no Kafka bootstrap brokers defined, please set the -brokers flag")
	}

	if len(topic) == 0 {
		panic("no topic given to be consumed, please set the -topic flag")
	}

}

func main() {
	// make a writer that produces to topic-A, using the round-robin distribution
	addrs := strings.Split(brokers, ",")
	w := &kafka.Writer{
		Addr:     kafka.TCP(addrs...),
		Topic:    topic,
		Balancer: &kafka.Hash{},
	}

	log.Print("Starting producer program...")
	log.Print(fmt.Sprintf("Brokers (%s), topic (%s)", brokers, topic))

	for {
		candidate_id := gofakeit.Number(1, 6)
		location := CandidateLocation{
			ID:  candidate_id,
			Lat: gofakeit.Latitude(),
			Lon: gofakeit.Longitude(),
		}

		msg, err := json.Marshal(location)

		if err != nil {
			panic(err)
		}

		payload := kafka.Message{
			Key:   []byte(strconv.Itoa(candidate_id)),
			Value: []byte(msg),
		}

		err = w.WriteMessages(context.Background(), payload)

		if err != nil {
			log.Fatal("failed to write messages:", err)
		}

		log.Print(fmt.Sprintf("message written at topic %s: %s = %s\n", w.Topic, string(payload.Key), string(payload.Value)))

		time.Sleep(1 * time.Second)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
