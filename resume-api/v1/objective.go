package v1

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mikegio27/resume-api/services"
)

func Objective(c *gin.Context) {
	objective, err := services.GetObjective()
	log.Println("Objective request received...")
	if err != nil {
		log.Println("Error retrieving objective: ", err)
		c.JSON(500, gin.H{"error": "Failed to fetch objective"})
		return
	}
	c.JSON(200, gin.H{"objective": objective})
}
