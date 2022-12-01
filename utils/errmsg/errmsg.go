package errmsg

const (
	SUCCSE = 200
	ERROR  = 500

	// 1000...  用户模块错误
	ERROR_USERNAME_USED      = 1001
	ERROR_PASSWORD_WRONG     = 1002
	ERROR_USERNAME_NOT_EXIST = 1003
	ERROR_TOKEN_NOT_EXIST    = 1004
	ERROR_TOKEN_RUNTIME      = 1005
	ERROR_TOKEN_WRONG        = 1006
	ERROR_TOKEN_TYPE_WRONG   = 1007
	ERROR_USER_NO_RIGHT      = 1008

	// 2000...  文章模块错误
	ERROR_ARTICLE_NOTEXIST = 2001
	// 3000...  分类模块错误
	ERROR_CATEGORYNAME_USED = 3001
	ERROR_CATEGORY_NOTEXIST = 3002
)

var codemsg = map[int]string{
	SUCCSE: "OK",
	ERROR:  "FAIL",

	ERROR_USERNAME_USED:      "用户名已存在！",
	ERROR_PASSWORD_WRONG:     "密码错误",
	ERROR_USERNAME_NOT_EXIST: "用户不存在",
	ERROR_TOKEN_NOT_EXIST:    "token不存在",
	ERROR_TOKEN_RUNTIME:      "token已过期",
	ERROR_TOKEN_WRONG:        "token错误",
	ERROR_TOKEN_TYPE_WRONG:   "token格式错误",
	ERROR_USER_NO_RIGHT:      "该用户无权限",

	ERROR_CATEGORYNAME_USED: "分类已存在",
	ERROR_CATEGORY_NOTEXIST: "分类不存在",

	ERROR_ARTICLE_NOTEXIST: "文章不存在",
}

// 获取错误信息函数
func GetErrmsg(code int) string {
	return codemsg[code]
}
