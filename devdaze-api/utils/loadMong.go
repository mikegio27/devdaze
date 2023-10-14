package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mikegio27/devdaze/devdaze-api/database"
)

func LoadMongoData(file_path string, db_name string, mongo_uri string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := database.Client.Database(db_name)

	fileContent, err := os.ReadFile(file_path)
	if err != nil {
		log.Fatalf("Failed to read the file: %v", err)
	}

	var resume map[string]interface{}
	if err := json.Unmarshal(fileContent, &resume); err != nil {
		log.Fatalf("Failed to unmarshal the JSON: %v", err)
	}

	for key, value := range resume {
		collection := db.Collection(key)
		fmt.Printf("Loading data into %s.%s...\n", db_name, key)

		_, err := collection.InsertOne(ctx, value)
		if err != nil {
			log.Fatalf("Failed to insert data into MongoDB: %v", err)
		}

		fmt.Printf("Loaded data from %s into %s.%s.\n", file_path, db_name, key)
	}
}
