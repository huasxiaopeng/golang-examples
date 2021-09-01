package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//路径中的请求参数
func main() {
	roy:=gin.Default()
	// This handler will match /user/john but will not match /user/ or /user
	roy.GET("/user/:name", func(c *gin.Context) {
		//应该上下文获取数据
		name:=c.Param("name")
		//响应体组装返回数据
		c.String(http.StatusOK,"hello %s",name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	//匹配补上的，这就是默认兜底策略进行匹配
	roy.GET("/user/:name/*action", func(c *gin.Context) {
		name:=c.Param("name")
		action:=c.Param("action")
		message:=name+" is"+action
		c.String(http.StatusOK,message)
	})

	//For each matched request Context will hold the route definition
	roy.POST("/user/:name/*action", func(c *gin.Context) {
		//检查路径是否匹配
		b:=c.FullPath() =="/user/:name/*action"
		c.String(http.StatusOK,"%t",b)

	})
	roy.Run(":8081")
}
