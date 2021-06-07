package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

// Cors 跨域全局中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		result, _ := regexp.MatchString("(.*sushao\\.top$|.*localhost:\\w+|[\\d.]+\\w+)", origin)
		if result {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-LabelIds, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-LabelIds")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		} else {
			fmt.Println(origin)
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			c.Next()
		}
	}
}
