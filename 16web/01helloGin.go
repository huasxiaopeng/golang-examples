package main
/**
   hello wold 程序
 */
import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)
func main() {
   r:=gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message":"hello lktbz",
		})
	})
   r.Run()
}
