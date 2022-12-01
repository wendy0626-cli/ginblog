// 文章模型

package model

import (
	"ginblog/utils/errmsg"

	"github.com/jinzhu/gorm"
)

// 文章模型
type Article struct {
	gorm.Model

	Category Category `gorm:"foreignkey:Cid"` //外键或者关联关系实现

	Title        string `gorm:"type:varchar(1024);not null" json:"title"`
	Flag         string `gorm:"type:varchar(20);not null" json:"flag"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(1024)" json:"desc"`
	Content      string `gorm:"type:longtext;not null" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

// 新增文章
func CreateArticle(data *Article) int {
	err := db.Create(data).Error

	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE
}

// GetArtInfo 查询单个文章
func GetOneArticle(id int) (Article, int) {
	var art Article
	err := db.Where("id = ?", id).Preload("Category").First(&art).Error
	db.Model(&art).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	if err != nil {
		return art, errmsg.ERROR_ARTICLE_NOTEXIST
	}
	return art, errmsg.SUCCSE
}

// 查询文章列表
func GetArticles(keywords string, pageSize int, pageNum int) ([]Article, int, int) {
	var articleList []Article
	var total int

	if keywords == "" {
		err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Preload("Category").Find(&articleList).Error
		// 单独计数
		db.Model(&articleList).Count(&total)
		if err != nil {
			return nil, errmsg.ERROR, 0
		}
		return articleList, errmsg.SUCCSE, total
	}

	err := db.Select("article.id,title, img, created_at, updated_at, `desc`, comment_count, read_count, category.name").Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Preload("Category").Where("title LIKE ?", keywords+"%").Find(&articleList).Error
	// 单独计数
	db.Model(&articleList).Where("title LIKE ?", keywords+"%").Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCSE, total

}

// @title	GetCateArt
// @description	查询分类某个分类下的文章
// @param	id	int 分类id
// @param	pageSize	int 分页大小
// @param	pageNum	int 页码
// @return	[]Article  article切片
// @return	int  错误代码
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int, int) {
	var cateArtList []Article
	var total int

	err := db.Preload("Category").Where("cid = ?", id).Find(&cateArtList).Count(&total).Limit(pageSize).Offset(pageSize * (pageNum - 1)).Error
	if err != nil {
		return nil, errmsg.ERROR_CATEGORY_NOTEXIST, 0
	}
	return cateArtList, errmsg.SUCCSE, total
}

// 编辑文章
func EditArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["flag"] = data.Flag
	maps["img"] = data.Img
	err := db.Model(&article).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除文章
func DeleteArticle(id int) int {
	var article Article
	err := db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
