package app

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hey I'm working")
	})
	router.POST("/", controllers.GetItem)

	router.Run(":8080")
}
