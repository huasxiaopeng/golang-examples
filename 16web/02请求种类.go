package main

import "github.com/gin-gonic/gin"

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}

func options(context *gin.Context) {
	
}

func head(context *gin.Context) {
	
}

func patching(context *gin.Context) {
	
}

func deleting(context *gin.Context) {
	
}

func putting(context *gin.Context) {
	
}

func posting(context *gin.Context) {

}

func getting(context *gin.Context) {
	
	
}
