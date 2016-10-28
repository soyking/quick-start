package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/soyking/quick-start/gin/errors"
)

func InitRouter(g *gin.Engine) {
	g.GET("/get", resultWrapper(func(c *gin.Context) (interface{}, error) {
		return nil, errors.ErrorParamsMissing
	}))
}
