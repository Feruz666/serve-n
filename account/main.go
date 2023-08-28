package main

import (
	"fmt"
	"log"

	account "github.com/Feruz666/serve-n/account/protos"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/protobuf/proto"
)

func main() {
	topic := "account-topic"
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "foo",
		"auto.offset.reset": "smallest",
	})

	if err != nil {
		log.Fatal(err)
	}

	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("I started \n")

	person := &account.Person{}

	for {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			err = proto.Unmarshal(e.Value, person)
			fmt.Printf("processing order: %#v\n", person)
		case kafka.Error:
			fmt.Printf("%v\n", e)
		}
	}
}
