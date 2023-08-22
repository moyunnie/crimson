package api

import (
	"crimson/service"
	"github.com/gin-gonic/gin"
)

type EditApi struct {
}

func (e *EditApi) RunEditDocker(c *gin.Context) {
	var dockerApi service.DockerApi
	containerid := dockerApi.CreateContainer()
	dockerApi.StartContainer(containerid)
	c.JSON(200, gin.H{"data": "你好"})
	return
}
