// user控制器

package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"ginblog/utils/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 定义返回消息状态代码
var code int

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var msg string

	_ = c.ShouldBindJSON(&data) //绑定前端数据
	msg, code = validator.Validate(&data)
	if code != errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	//查询用户是否存在
	code = model.CheckUser(data.Username)
	// 用户已存在
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	// 用户不存在
	if code == errmsg.SUCCSE {
		// 执行创建新用户
		model.CreateUser(&data)
	}

	c.JSON(http.StatusOK, gin.H{ //状态码及响应头
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrmsg(code),
	})

}

//查询单个用户
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetUserInfo(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrmsg(code),
	})

}

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, total := model.GetUsers(username, pageSize, pageNum)
	code = errmsg.SUCCSE

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrmsg(code),
	})
}

// 编辑用户信息 更新
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckUpUser(id, data.Username)
	// 用户已存在
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	// 用户不存在，执行更新
	if code == errmsg.SUCCSE {
		model.EditUser(id, &data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrmsg(code),
	})
}

// 重置密码
func ResetPwd(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.ResetPwd(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrmsg(code),
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrmsg(code),
	})
}
