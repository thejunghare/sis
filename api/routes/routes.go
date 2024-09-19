package routes

import (
	"erp/api/api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Route() {
	r := gin.Default()

	r.Use(cors.Default())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.POST("/login", controllers.Login)

	r.Run()
}
