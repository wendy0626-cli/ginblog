package main

import (
	"ginblog/model"
	"ginblog/routes"
)

func main() {
	// 引用数据库
	model.InitDb()

	// 运行路由
	routes.InitRouter()
}
