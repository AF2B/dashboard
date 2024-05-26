package repository

import (
	"context"
	"dashboard/internal/models"
)

func CreateActivityVisitor(visitor *models.ActivityVisitor) error {
	config := MongoDBConfig{
		URI: "mongodb://localhost:27017",
		DBName: "analytics-dashboard",
	}

	conn := MongoDBConnection{config: config}
	err := conn.Connect()
	if err != nil {
		return err
	}
	defer conn.client.Disconnect(context.Background())

	collection := conn.client.Database(config.DBName).Collection("activity_visitors")
	_, err = collection.InsertOne(context.Background(), visitor)
	if err != nil {
		return err
	}

	return nil
}