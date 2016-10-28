package routers

import "github.com/gin-gonic/gin"

type resultHandler func(*gin.Context) (interface{}, error)

func resultWrapper(handler resultHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, buildResponse(handler(c)))
	}
}

type noResultHandler func(*gin.Context) error

func noResultWrapper(handler noResultHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, buildResponse(nil, handler(c)))
	}
}

type pageHandler func(*gin.Context) (int, interface{}, error)

func pageWrapper(handler pageHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, buildPageResponse(handler(c)))
	}
}
