package v1

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Health check
//	@Summary	Checks if the server is alive
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}	"OK"
func Healthcheck(c *gin.Context) {
	log.Println("Received Healthcheck")
	c.IndentedJSON(200, "OK")
}
