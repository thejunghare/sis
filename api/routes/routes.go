package routes

import (
	"erp/api/api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Route() {
	port := "8080"
	r := gin.Default()

	// fix cors
	r.Use(cors.Default())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// for session
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// for serve react
	r.Static("/static", "./build/static")

	r.POST("/login", controllers.Login)

	r.GET("/dashboard", func(c *gin.Context) {
		c.File("./build/index.html")
	})

	r.Run(":" + port)
}
