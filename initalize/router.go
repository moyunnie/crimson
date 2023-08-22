package initalize

import (
	"crimson/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	c := gin.Default()
	rou := router.InitRouter(c)
	err := rou.Run(":8012")
	if err != nil {
		panic(err)
	}
}
