package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func New(mongoClient *mongo.Client) Models {
	client = mongoClient
	return Models{
		LogEntry: LogEntry{},
	}
}

type Models struct {
	LogEntry LogEntry
}

type LogEntry struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `bson:"name" json:"name"`
	Data      string    `bson:"data" json:"data"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func (l *LogEntry) Insert(entry LogEntry) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), LogEntry{
		Name:      entry.Name,
		Data:      entry.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Println("Error inserting log entry:", err)
		return err
	}
	return nil
}

func (l *LogEntry) All() ([]*LogEntry, error) {
	collection := client.Database("logs").Collection("logs")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSort(map[string]int{"_id": -1})
	cur, err := collection.Find(context.TODO(), bson.D{}, options)
	if err != nil {
		log.Println("Error getting all log entries:", err)
		return nil, err
	}
	defer cur.Close(ctx)

	var logs []*LogEntry
	for cur.Next(ctx) {
		var item LogEntry
		err := cur.Decode(&item)
		if err != nil {
			log.Println("Error decoding log entry:", err)
			return nil, err
		} else {
			logs = append(logs, &item)
		}
		logs = append(logs, &item)
	}

	return logs, nil
}

func (l *LogEntry) GetDetails(id string) (*LogEntry, error) {
	collection := client.Database("logs").Collection("logs")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error getting log entry by id:", err)
		return nil, err
	}

	var logEntry LogEntry
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&logEntry)
	if err != nil {
		log.Println("Error getting log entry by id:", err)
		return nil, err
	}

	return &logEntry, nil
}

func (l *LogEntry) DropCollection() error {
	collection := client.Database("logs").Collection("logs")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := collection.Drop(ctx)
	if err != nil {
		log.Println("Error dropping log collection:", err)
		return err
	}
	return nil
}

func (l *LogEntry) Update() (*mongo.UpdateResult, error) {
	collection := client.Database("logs").Collection("logs")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	docID, err := primitive.ObjectIDFromHex(l.ID)
	if err != nil {
		log.Println("Error updating log entry:", err)
		return nil, err
	}

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.D{
			{"$set", bson.D{
				{"name", l.Name},
				{"data", l.Data},
				{"updated_at", time.Now()},
			}},
		})
	if err != nil {
		log.Println("Error updating log entry:", err)
		return nil, err
	}

	return result, nil
}
