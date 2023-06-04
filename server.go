package main

import (
	"github.com/gin-gonic/gin"
	"go-social-network/server/controllers"
	"go-social-network/server/middlewares"
	"go-social-network/server/models"
	"log"
)

func main() {
	models.ConnectToDb()
	r := gin.Default()
	public := r.Group("/api/public/v1")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/private/v1").Use(middlewares.JwtAuthMiddleware())
	protected = protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/me", controllers.CurrentUser)

	if err := r.Run(":8080"); err != nil {
		log.Fatalln("Error running server ", err)
	}
}
