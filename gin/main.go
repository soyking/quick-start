package main

import (
	"github.com/gin-gonic/gin"
	"github.com/soyking/quick-start/gin/config"
	"github.com/soyking/quick-start/gin/routers"
)

func main() {
	c := config.InitConfig()
	g := gin.Default()
	routers.InitRouter(g)
	g.Run(":" + c.Port)
}
