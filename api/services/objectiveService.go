package services

import (
	"context"
	"log"

	"github.com/mikegio27/api/database"
	"go.mongodb.org/mongo-driver/bson"
)

func GetObjective() (string, error) {
	var result struct {
		Objective string `bson:"objective"`
	}

	collection := database.Client.Database("resume").Collection("resume")
	err := collection.FindOne(context.TODO(), bson.D{}).Decode(&result)

	if err != nil {
		log.Println("Error retrieving objective:", err)
		return "", err
	}

	return result.Objective, nil
}
