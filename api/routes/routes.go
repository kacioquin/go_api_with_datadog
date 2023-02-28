package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kacioquin/go_api_with_datadog/api/controllers"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	tweetController := controllers.NewTweetController()

	v1 := router.Group("/v1")
	{
		v1.GET("/tweets", tweetController.FindAll)
		v1.POST("/tweet", tweetController.Create)
		v1.GET("/tweet/:id", tweetController.Delete)
	}

	return v1
}