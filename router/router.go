package router

import (
	"crimson/api"
	"github.com/gin-gonic/gin"
)

func InitRouter(c *gin.Engine) *gin.Engine {
	var controller api.EditApi
	v1 := c.Group("edit")
	v1.POST("run", controller.RunEditDocker)
	return c
}
