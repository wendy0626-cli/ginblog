package model

import (
	"ginblog/utils/errmsg"
)

type Profile struct {
	ID     int    `gorm:"primarykey;autoIncrement" json:"id"`
	Name   string `gorm:"type:varchar(50)" json:"name"`
	Motto  string `gorm:"type:varchar(255)" json:"motto"`
	Desc   string `gorm:"type:varchar(1024)" json:"desc"`
	Qqchat string `gorm:"type:varchar(200)" json:"qq_chat"`
	Wechat string `gorm:"type:varchar(200)" json:"wechat"`
	Email  string `gorm:"type:varchar(200)" json:"email"`
	Gitee  string `gorm:"type:varchar(1024)" json:"gitee"`
	Bili   string `gorm:"type:varchar(1024)" json:"bili"`
	Img    string `gorm:"type:varchar(1024)" json:"img"`
	Avatar string `gorm:"type:varchar(1024)" json:"avatar"`
}

// 获取个人信息
func GetProfile(id int) (Profile, int) {
	var profile Profile

	err := db.Where("ID = ?", id).First(&profile).Error
	if err != nil {
		return profile, errmsg.ERROR
	}
	return profile, errmsg.SUCCSE
}

// 更新个人信息
func UpdateProfile(id int, data *Profile) int {
	var profile Profile

	err := db.Model(&profile).Where("ID = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
