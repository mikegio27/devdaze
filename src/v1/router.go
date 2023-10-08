package v1

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/resume")
	{
		api.GET("/experience", Experience)
		api.GET("/objective", Objective)
		api.GET("/techskills", TechSkills)
	}
	return r
}
