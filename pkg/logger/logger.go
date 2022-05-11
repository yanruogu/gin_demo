package logger

import (
	"fmt"

	"github.com/natefinch/lumberjack"
	"github.com/yanruogu/gin_demo/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	LoggerHandler *zap.Logger
	cfg           *config.Log
}

func New(cfg *config.Log) *Logger {
	return &Logger{
		cfg: cfg,
	}
}

func (l *Logger) Init() error {
	encoder := l.getLogEncoder()
	writeSyncer := l.getLogWriter()
	// zapcore.DebugLevel定义输出的日志级别为debug级别
	level, err := l.getLogLevel()
	if err != nil {
		panic(fmt.Errorf("init log failed, err: %s", err))
	}
	core := zapcore.NewCore(encoder, writeSyncer, level)
	l.LoggerHandler = zap.New(core)
	zap.ReplaceGlobals(l.LoggerHandler)
	//zap.L().Info("init logger success!")
	return nil
}

//var logger *zap.Logger

// 定义日志格式
func (l *Logger) getLogEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 定义输出的时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 将日志级别设置为大写
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 定义输出位置
func (l *Logger) getLogWriter() zapcore.WriteSyncer {
	//filename, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   l.cfg.Path,
		MaxSize:    l.cfg.MaxSize,
		MaxBackups: l.cfg.MaxBackups,
		MaxAge:     l.cfg.MaxAge,
		Compress:   false,
	}

	return zapcore.AddSync(lumberJackLogger)
}

// 定义输出的日志级别
func (l *Logger) getLogLevel() (*zapcore.Level, error) {
	var level = new(zapcore.Level)
	err := level.UnmarshalText([]byte(l.cfg.Level))
	if err != nil {
		return nil, err
	}
	return level, nil
}
