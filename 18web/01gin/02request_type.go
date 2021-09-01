package main

import "github.com/gin-gonic/gin"

//get ,post ,delete
func main() {
  router:= gin.Default()

  //router.GET()
	router.GET("/someGet")
	router.POST("/somePost")
	router.PUT("/somePut")
	router.DELETE("/someDelete")
	router.PATCH("/somePatch")
	router.HEAD("/someHead")
	router.OPTIONS("/someOptions")
	//指定地址启动
    router.Run(":8081")

}
