package v1

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Healthcheck(c *gin.Context) {
	log.Println("Received Healthcheck")
	c.IndentedJSON(200, "OK")
}
