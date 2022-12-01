package validator

import (
	"fmt"
	"ginblog/utils/errmsg"
	"github.com/go-playground/locales/zh_Hans_CN"
	untrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

//数据验证
func Validate(data interface{}) (string, int) {
	validate := validator.New()                 //实例化验证器
	uni := untrans.New(zh_Hans_CN.New())        //指定翻译为中文
	trans, _ := uni.GetTranslator("zh_Hans_CN") //创建翻译器

	err := zhTrans.RegisterDefaultTranslations(validate, trans) //注册验证翻译
	if err != nil {
		fmt.Println("err:", err)
	}
	//进行映射
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data) //进行验证（结构体）
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) { //遍历错误信息
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCSE
}
