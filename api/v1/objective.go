package v1

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mikegio27/devdaze/api/services"
)

// Objective godoc
//
//	@Summary	Resume objective
//	@Produce	json
//
// @Success 200 {object} map[string]string "Objective"
// @Router /objective [get]
func Objective(c *gin.Context) {
	objective, err := services.GetObjective()
	if err != nil {
		log.Println("Error retrieving objective: ", err)
		c.JSON(500, gin.H{"error": "Failed to fetch objective"})
		return
	}
	c.JSON(200, gin.H{"objective": objective})
}
