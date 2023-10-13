package v1

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mikegio27/resume-api/services"
)

// Tech Skills
//
//	@Summary	Resume tech skills summary
//	@Produce	json
//	@Success	200	{array} 	services.TechSkills
//
// @Router /tech-skills [get]
func TechSkills(c *gin.Context) {
	techSkills, err := services.GetTechSkills()
	if err != nil {
		log.Println("Error retrieving tech skills: ", err)
		c.JSON(500, "Error: Failed to retrieve Tech Skills.")
		return
	}
	c.IndentedJSON(200, techSkills)
}
