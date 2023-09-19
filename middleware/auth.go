package middleware

import (
	"strings"

	"github.com/TravisRoad/travis-blog-service/global"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		tokenSlice := strings.Split(authorization, " ")
		if len(tokenSlice) < 2 {
			c.AbortWithStatus(401)
			return
		}
		if tokenSlice[1] != global.Config.Token {
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}
