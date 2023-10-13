package v1

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mikegio27/resume-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://devdaze.org/", "https://devdaze.org", "https://api.devdaze.org", "http://localhost:3000"}
	config.AllowMethods = []string{"PUT", "GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	log.Println("Router started...")
	r.Use(cors.New(config))
	r.GET("/health", Healthcheck)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/resume")
	{
		api.GET("/experience", Experience)
		api.GET("/objective", Objective)
		api.GET("/techskills", TechSkills)
	}
	return r
}
