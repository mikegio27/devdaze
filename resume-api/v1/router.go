package v1

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"PUT", "GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	log.Println("Router started...")
	r.Use(cors.New(config))
	r.GET("/health", Healthcheck)
	api := r.Group("/resume")
	{
		api.GET("/experience", Experience)
		api.GET("/objective", Objective)
		api.GET("/techskills", TechSkills)
	}
	return r
}
