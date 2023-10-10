package v1

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/resume")
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.MaxAge = 12 * time.Hour

	api.Use(cors.New(config))
	{
		api.GET("/experience", Experience)
		api.GET("/objective", Objective)
		api.GET("/techskills", TechSkills)
	}
	return r
}
