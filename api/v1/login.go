package v1

import (
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	var token string
	var code int

	c.ShouldBindJSON(&data)

	code = model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCSE {
		token, code = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrmsg(code),
		"token":   token,
	})
}
