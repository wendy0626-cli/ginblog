package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 路由初始函数
func InitRouter() {
	// 设置AppMode（app运行模式）
	gin.SetMode(utils.AppMode)

	// 初始化gin
	r := gin.New()

	//启用跨域中间件
	r.Use(middleware.Cors())
	//启用日志中间件
	r.Use(middleware.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// 托管静态网页 --------------------------------------------------------------------
	// 加载html页面
	r.LoadHTMLFiles("public/admin/admin.html", "public/front/front.html")
	// r.LoadHTMLFiles("web/admin/dist/index.html", "web/front/dist/index.html")
	// 后台静态资源
	r.Static("/admin/static", "public/admin/static")
	// 前台静态资源
	r.Static("/front/static", "public/front/static")
	// 图标
	// r.StaticFile("/favicon.ico", "public/favicon.ico")
	r.StaticFile("/favicon.ico", "public/favicon.ico")
	// 后台路由
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", nil)
	})
	// 前台路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "front.html", nil)
	})
	// -------------------------------------------------------------------------------

	// 需要权限的路由组
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken()) //使用jwt中间件
	{
		// 用户模块的路由接口
		auth.POST("user/add", v1.AddUser)       // 添加用户
		auth.PUT("user/:id", v1.EditUser)       // 编辑用户
		auth.PUT("user/:id/reset", v1.ResetPwd) // 重置密码
		auth.DELETE("user/:id", v1.DeleteUser)  // 删除用户

		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)      // 添加分类
		auth.PUT("category/:id", v1.EditCategory)      // 编辑分类
		auth.DELETE("category/:id", v1.DeleteCategory) // 删除分类

		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)      // 添加文章
		auth.PUT("article/:id", v1.EditArticle)      // 编辑文章
		auth.DELETE("article/:id", v1.DeleteArticle) //删除文章

		// 个人信息路由
		auth.PUT("profile/:id", v1.UpdateProfile) //更新个人信息

		// 上传文件
		auth.POST("upload", v1.Upload)
		// 评论模块
		auth.GET("comment/list", v1.GetCommentList)
		auth.DELETE("delcomment/:id", v1.DeleteComment)
		auth.PUT("checkcomment/:id", v1.CheckComment)
		auth.PUT("uncheckcomment/:id", v1.UncheckComment)

	}

	// 不需授权的路由组
	routerV1 := r.Group("api/v1")
	{
		// 登录
		routerV1.POST("login", v1.Login)
		// 用户路由
		routerV1.GET("user/:id", v1.GetUserInfo) // 获取单个用户
		routerV1.GET("users", v1.GetUsers)       // 获取用户列表
		// 分类路由
		routerV1.GET("category/:id", v1.GetCategory)        // 获取单个列表
		routerV1.GET("categories", v1.GetCategories)        // 获取分类列表
		routerV1.GET("category/:id/artlist", v1.GetCateArt) //获取分类下所有文章
		// 文章路由
		routerV1.GET("article/:id", v1.GetArticle) // 查询单个文章
		routerV1.GET("articles", v1.GetArticles)   // 获取文章列表
		// 个人信息路由
		routerV1.GET("profile/:id", v1.GetProfile) //获取个人信息
		// 评论模块
		routerV1.POST("addcomment", v1.AddComment)
		routerV1.GET("comment/info/:id", v1.GetComment)
		routerV1.GET("commentfront/:id", v1.GetCommentListFront)
		routerV1.GET("commentcount/:id", v1.GetCommentCount)
	}

	// 运行gin
	//_ = r.Run(utils.HttpPort)

	// 运行的地址和端口
	_ = r.Run((utils.HttpAddress + utils.HttpPort))
}
