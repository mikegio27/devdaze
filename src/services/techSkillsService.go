package services

import (
	"context"
	"log"

	"github.com/mikegio27/resume-api/src/database"
	"go.mongodb.org/mongo-driver/bson"
)

type TechSkills struct {
	Languages  []string `bson:"languages"`
	Platforms  []string `bson:"platforms"`
	Linux      []string `bson:"linux"`
	CICD       []string `bson:"cicd"`
	Monitoring []string `bson:"monitoring"`
	IAC        []string `bson:"iac"`
	Databases  []string `bson:"databases"`
}

func GetTechSkills() (TechSkills, error) {
	var result struct {
		TechSkills TechSkills `bson:"techSkills"`
	}

	collection := database.Client.Database("resume").Collection("resume")
	err := collection.FindOne(context.TODO(), bson.D{}).Decode(&result)

	if err != nil {
		log.Println("Error retrieving tech skills:", err)
		return TechSkills{}, err
	}

	return result.TechSkills, nil
}
