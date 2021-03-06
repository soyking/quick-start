package main

import (
	"github.com/gin-gonic/gin"
	"{{.package}}/config"
	"{{.package}}/routers"
)

func main() {
	c := config.InitConfig()
	g := gin.Default()
	routers.InitRouter(g)
	g.Run(":" + c.Port)
}
