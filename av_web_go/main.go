package main

import (
	"av_process/av_web_go/pkg/config"
	"av_process/common"
	"time"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

// 4a081c2dd82c66a948f292873071f4fd
func main() {
	loggerConfig := &common.LoggerConfig{
		Level:        logrus.DebugLevel,
		MaxAge:       time.Hour * 24 * 7,
		RotationTime: time.Hour * 24,
		LogPath:      "./log",
		LogFileName:  "av_web.log",
		ServiceName:  "av_web",
		Version:      "v0.1.0",
	}

	loggerConfig.NewLogger()

	logger = common.Logger
	logger.Info("start")

	conf, err := config.NewConfig()
	if err != nil {
		logger.Error("读取配置文件失败： ", err)
		return
	}

	router := createRouter()
	router.Run(conf.Web["addr"])
}
