package v1

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mikegio27/resume-api/services"
)

func Experience(c *gin.Context) {
	experience, err := services.GetExperience()
	log.Println("Received request for experience...")
	if err != nil {
		log.Println("Error retrieving experience: ", err)
		c.JSON(500, gin.H{"error": "Failed to fetch experience"})
		return
	}
	c.IndentedJSON(200, experience)
}
