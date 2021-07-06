package middlewares

import (
	"gin-data/app/services"
	jwt2 "gin-data/utils/jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	lichv "github.com/lichv/go"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := c.DefaultQuery("token", "")
		if token == "" {
			token = c.DefaultPostForm("token", "")
		}
		if token == "" {
			token, _ = c.Cookie("token")
		}
		if token == "" {
			token = c.GetHeader("X-TOKEN")
		}

		if token == "" {
			code = 400
		} else {
			claims, err := jwt2.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = 404
				default:
					code = 405
				}
			} else {
				list, _ := services.GetAllWhitelistUserCode(map[string]interface{}{}, "code asc", -1)
				if lichv.In(list,claims.Code) {
					c.JSON(http.StatusUnauthorized, gin.H{
						"code": code,
						"msg":  "不在白名单中，请联系管理员",
						"data": claims,
					})
					c.Abort()
					return
				}
			}
		}

		if code != 200 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "token失效",
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
