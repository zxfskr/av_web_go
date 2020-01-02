package common

import (
	"path"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// Logger 日志
var Logger *logrus.Entry

// LoggerConfig logger的配置
type LoggerConfig struct {
	Level         logrus.Level
	MaxAge        time.Duration
	RotationTime  time.Duration
	LogPath       string
	LogFileName   string
	contextLogger *logrus.Logger
	ServiceName   string
	Version       string
}

func (s *LoggerConfig) configureLogToLocal() {
	s.contextLogger = logrus.New()
	baseLogPath := path.Join(s.LogPath, s.LogFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d",
		rotatelogs.WithLinkName(baseLogPath),        // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(s.MaxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(s.RotationTime), // 日志切割时间间隔
	)
	if err != nil {
		logrus.Fatalf("配置日志本地化操作失败, 错误原因: %v", err)
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableTimestamp: false,
	})
	s.contextLogger.AddHook(lfHook)
	s.contextLogger.SetLevel(s.Level)
}

// NewLogger 新建一个logger
func (s *LoggerConfig) NewLogger() {
	s.configureLogToLocal()
	s.contextLogger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableTimestamp: false,
	})

	Logger = s.contextLogger.WithFields(logrus.Fields{
		"serviceName": s.ServiceName,
		"machineIP":   GetLocalIPAddress(),
		"version":     s.Version,
	})

}
