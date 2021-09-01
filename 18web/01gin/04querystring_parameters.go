package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/welcome", func(c *gin.Context) {
		//带兜底策略的参数
		firstname:=c.DefaultQuery("firstname","guest")
		//没有兜底策略
		lastname:=c.Query("lastname")
		c.String(http.StatusOK,"hello %s %s",firstname,lastname)
	})
	router.Run(":8081")
}
