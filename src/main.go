package main

import (
	"github.com/mikegio27/resume-api/database"
	v1 "github.com/mikegio27/resume-api/v1"
)

func main() {
	database.ConnectMongoDB("mongodb://admin:password@localhost:27017/")
	r := v1.SetupRouter()

	r.Run(":8080")
}
