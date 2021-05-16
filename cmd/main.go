package main

import (
	"opiece/server/core"
	"opiece/server/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	core.RunServer()
}

func ginRun() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	authMiddleware := middleware.NewJWTMiddleware()
	r.POST("/api/login", authMiddleware.LoginHandler)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	auth := r.Group("/api/auth")
	auth.POST("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
