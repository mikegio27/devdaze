package main

import (
	"fmt"
	"os"

	"github.com/mikegio27/resume-api/database"
	v1 "github.com/mikegio27/resume-api/v1"
)

func main() {
	username := os.Getenv("MONGO_USERNAME")
	pass := os.Getenv("MONGO_PASSWORD")
	uri := fmt.Sprintf("mongodb://%s:%s@localhost:27017", username, pass)
	//uri := fmt.Sprintf("mongodb://%s:%s@web-apps-mongodb.web-apps.svc.cluster.local:27017/resume", username, pass)
	database.ConnectMongoDB(uri)
	r := v1.SetupRouter()

	r.Run(":8080")
}
