package services

import (
	"context"
	"log"

	"github.com/mikegio27/api/database"
	"go.mongodb.org/mongo-driver/bson"
)

type Experience struct {
	Title      string   `bson:"title"`
	Company    string   `bson:"company"`
	Department string   `bson:"department"`
	Date       string   `bson:"date"`
	Duties     []string `bson:"duties"`
}

func GetExperience() ([]Experience, error) {
	var result struct {
		Experience []Experience `bson:"experience"`
	}

	collection := database.Client.Database("resume").Collection("resume")
	err := collection.FindOne(context.TODO(), bson.D{}).Decode(&result)

	if err != nil {
		log.Println("Error retrieving experience:", err)
		return nil, err
	}

	return result.Experience, nil
}
