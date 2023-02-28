package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kacioquin/go_api_with_datadog/api/routes"
)

func main() {
	app := gin.Default()

	routes.AppRoutes(app)

	app.Run(":8081")
}
