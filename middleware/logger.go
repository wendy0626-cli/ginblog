package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	fielPath := "log/ginblog"
	linkName := "latast_log.log"
	src, err := os.OpenFile(fielPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err:", err)
	}

	logger := logrus.New()

	logger.Out = src                   //日志输出文件
	logger.SetLevel(logrus.DebugLevel) //日志级别

	//日志分割
	logWriter, _ := rotalog.New(
		fielPath+"%Y%m%d.log",
		rotalog.WithMaxAge(7*24*time.Hour),
		rotalog.WithRotationTime(24*time.Hour),
		rotalog.WithLinkName(linkName),
	)

	//日志写出配置
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		// 时间格式化
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0))) //花费时间

		hostName, err := os.Hostname() //主机名
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()    //状态码
		clientIp := c.ClientIP()           //客户端IP
		userAgent := c.Request.UserAgent() //user agent
		dataSize := c.Writer.Size()        //数据大小
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method   //请求方法
		path := c.Request.RequestURI //请求地址

		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}

}
