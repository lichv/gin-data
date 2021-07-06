package middlewares

import (
	"gin-data/app/services"
	"gin-data/utils"
	"github.com/gin-gonic/gin"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		input := utils.GetMapFromContext(c)
		services.LogFromContext(c, path, "", input)
		c.Next()
	}
}
