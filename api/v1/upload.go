package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

//上传文件
func Upload(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")

	fielSize := fileHeader.Size

	url, code := model.UploadFile(file, fielSize)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrmsg(code),
		"url":     url,
	})
}
