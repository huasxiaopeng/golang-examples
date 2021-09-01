package main

import (
	"github.com/gin-gonic/gin"

	_ "net/http"
)
//url 路径绑定
type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}
func main() {
	router:=gin.Default()
	router.POST("/login", func(c *gin.Context) {
		var from LoginForm
		 if c.ShouldBind(&from)==nil{
		 	if from.User=="user"&& from.Password=="password"{
		 		c.JSON(200,gin.H{"status":"you are logged in"})
			}else{
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		 }
	})
	router.Run(":8081")
}

//curl -v --form user=user --form password=password http://localhost:8081/login
