package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/listener-service/event"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// This is a placeholder main function.
	// The actual implementation would go here.
	// try to connect to rabbitmq\
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()
	log.Println("Successfully connected to RabbitMQ")

	// start listening for messages

	// create consumer
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	// watch the queue and consume messages
	err = consumer.Listen([]string{"log.INFO", "log.ERROR", "log.WARNING"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {
	// connection string
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// don't continue until rabbitmq is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready")
			counts++
		} else {
			connection = c
			fmt.Println("Connected to RabbitMQ!")
			break
		}

		if counts >= 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		fmt.Println("Backing off for ", backOff)
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}
