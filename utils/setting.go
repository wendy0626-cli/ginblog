package utils

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

var (
	AppMode     string
	HttpAddress string
	HttpPort    string
	JwtKey      string

	Db     string
	DbHost string
	DbPort string
	DbUser string
	DbPwd  string
	DbName string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
)

// 配置初始化函数
func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取失败，请检查文件路径", err)
		os.Exit(1)
	}
	LoadServer(file)
	LoadData(file)
	LoadQiniu(file)
}

// 读取server节点配置
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpAddress = file.Section("server").Key("HttpAddress").MustString("0.0.0.0")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":5000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("4ti7ng2y0u")
}

// 读取database节点配置
func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3366")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPwd = file.Section("database").Key("DbPwd").MustString("Admin1234!")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

// 读取七牛云节点配置
func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuServer = file.Section("qiniu").Key("QiniuServer").String()
}
