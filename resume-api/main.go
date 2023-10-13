package main

import (
	"fmt"
	"os"

	"github.com/mikegio27/resume-api/database"
	_ "github.com/mikegio27/resume-api/docs"
	v1 "github.com/mikegio27/resume-api/v1"
)

// @title			DevDaze API
// @version		1.0
// @description	API currently serving my resume
// @host			api.devdaze.org
func main() {
	username := os.Getenv("MONGO_USERNAME")
	pass := os.Getenv("MONGO_PASSWORD")
	uri := fmt.Sprintf("mongodb://%s:%s@localhost:27017", username, pass)
	//uri := fmt.Sprintf("mongodb://%s:%s@web-apps-mongodb.web-apps.svc.cluster.local:27017/resume", username, pass)
	database.ConnectMongoDB(uri)
	//utils.LoadMongoData("resume.json", "resume", uri)
	r := v1.SetupRouter()

	r.Run(":8080")
}
