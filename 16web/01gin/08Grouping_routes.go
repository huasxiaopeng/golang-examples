package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login")
		v1.POST("/submit")
		v1.POST("/read")
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login")
		v2.POST("/submit")
		v2.POST("/read")
	}

	router.Run(":8080")
}
