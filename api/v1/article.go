// article控制器

package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article

	_ = c.ShouldBindJSON(&data)

	model.CreateArticle(&data)

	c.JSON(http.StatusOK, gin.H{ //状态码及响应头
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrmsg(code),
	})

}

// 查询单个文章
func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, code := model.GetOneArticle(id)

	// 把markdown格式文章内容转换未html
	mddata := []byte(article.Content)
	article.Content = string(template.HTML(blackfriday.Run(mddata)))

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errmsg.GetErrmsg(code),
	})
}

// 查询文章列表
func GetArticles(c *gin.Context) {
	keywords := c.Query("keywords")
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, code, total := model.GetArticles(keywords, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrmsg(code),
	})
}

// 查询分类下所有文章
func GetCateArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, code, total := model.GetCateArt(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrmsg(code),
	})
}

// 编辑文章
func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)

	model.EditArticle(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrmsg(code),
	})
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrmsg(code),
	})
}
