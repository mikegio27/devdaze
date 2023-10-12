package v1

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/resume")
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	api.Use(cors.New(config))
	{
		api.GET("/experience", Experience)
		api.GET("/objective", Objective)
		api.GET("/techskills", TechSkills)
	}
	return r
}
