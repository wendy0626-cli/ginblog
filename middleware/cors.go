package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowAllOrigins: true,
			//AllowOrigins:     []string{"*"},

			AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE", "UPDATE"},
			AllowHeaders:     []string{"*"},
			ExposeHeaders:    []string{"Content-Length", "Content-Type", "Authorization"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		})

	//return func(c *gin.Context) {
	//	method := c.Request.Method
	//
	//	c.Header("Access-Control-Allow-Origin", "*")
	//	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	//	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE, UPDATE")
	//	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	//	c.Header("Access-Control-Allow-Credentials", "true")
	//
	//	//放行所有OPTIONS方法
	//	if method == "OPTIONS" {
	//		c.AbortWithStatus(http.StatusNoContent)
	//	}
	//	// 处理请求
	//	c.Next()
	//}
}
