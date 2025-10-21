package main

import (
	"context"
	"log"
	"time"

	"github.com/logger-service/data"
)

type RPCServer struct{}

type RPCPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println("Error inserting log entry:", err)
		return err
	}
	*resp = "Processed payload via RPC: " + payload.Name
	return nil
}
