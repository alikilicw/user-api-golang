package main

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) UrlPatterns(router *gin.Engine) {
	api := router.Group("/api")

	authGroup := api.Group("/auth")
	authGroup.POST("/register", server.Register)
	authGroup.GET("/users/:id", server.GetUsers)
	authGroup.POST("/login", server.Login)
}
