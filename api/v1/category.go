// 分类控制器

package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category

	_ = c.ShouldBindJSON(&data)

	code = model.CheckCategory(data.Name)
	// 分类已存在
	if code == errmsg.ERROR_CATEGORYNAME_USED {
		code = errmsg.ERROR_CATEGORYNAME_USED
	}
	// 分类不存在
	if code == errmsg.SUCCSE {
		// 执行创建新分类
		model.CreateCategory(&data)
	}

	c.JSON(http.StatusOK, gin.H{ //状态码及响应头
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrmsg(code),
	})

}

// 获取一个分类
func GetCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetCategory(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrmsg(code),
	})
}

// 查询分类列表
func GetCategories(c *gin.Context) {
	name := c.Query("name")
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, total := model.GetCategories(name, pageSize, pageNum)
	code = errmsg.SUCCSE

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrmsg(code),
	})
}

// 编辑分类
func EditCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	// 分类已存在
	if code == errmsg.ERROR_CATEGORYNAME_USED {
		c.Abort()
	}
	// 分类不存在，执行更新
	if code == errmsg.SUCCSE {
		model.EditCategory(id, &data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrmsg(code),
	})
}

// 删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrmsg(code),
	})
}
