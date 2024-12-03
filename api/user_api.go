package api

import (
	"github.com/gin-gonic/gin"
	"sprout/service"
)

func Start() {
	r := gin.Default()
	r.GET("/user/getInfo", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
	r.POST("/user/saveInfo", func(context *gin.Context) {
		service.SaveUserInfo(context)
		context.JSON(200, gin.H{"message": "Hello, World!"})
	})
	r.Run("0.0.0.0:9999")
}
