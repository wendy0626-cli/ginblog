package model

import (
	"fmt"
	"ginblog/utils"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var dberr error

// 数据库初始函数
func InitDb() {
	// db,err := gorm.Open("mysql","root:123456@tcp(127.0.0.1:3366)/ginblog?charset=utf8&parseTime=True&loc=Local")
	// 连接数据库
	db, dberr = gorm.Open(utils.Db,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			utils.DbUser,
			utils.DbPwd,
			utils.DbHost,
			utils.DbPort,
			utils.DbName))

	if dberr != nil {
		fmt.Println("连接数据库失败，请检查参数！", dberr)
	}

	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	db.SingularTable(true)

	// 自动迁移
	db.AutoMigrate(&User{}, &Category{}, &Article{}, &Profile{}, Comment{})

	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(10 * time.Second)

	//db.Close()
}
