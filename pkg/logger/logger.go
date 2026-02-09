package logger

import (
	"fmt"
	"mygo_bangforai/pkg/config"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

func InitLogger() {
	logLevel:=getLogLevel()
	logFile:=config.GetLoggerConfig().File
	//logErrorFile:=config.GetLoggerConfig().ErrorFile
	logOutput:=config.GetLoggerConfig().Output
	logFormat:=config.GetLoggerConfig().Format

	encoderConfig := zap.NewProductionEncoderConfig()// 生产环境编码器配置
	encoderConfig.TimeKey = "time"					// 时间键名
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder// 时间编码器为ISO8601格式

	var encoder zapcore.Encoder
	if logFormat == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)// JSON编码器
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)// 控制台编码器
	}

	var core zapcore.Core
	if logOutput == "file" {
		file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("打开日志文件失败: %v\n", err)
			// 回退到控制台输出
			core = zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), logLevel)
		} else {
			// 同时输出到文件和控制台
			core = zapcore.NewTee(
				zapcore.NewCore(encoder, zapcore.AddSync(file), logLevel),
				zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), logLevel),
			)
		}
	} else {
		core = zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), logLevel)
	}

	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	Sugar = Logger.Sugar()
	
	Logger.Info("日志初始化完成",
		zap.String("level", logLevel.String()),
		zap.String("format", logFormat),
		zap.String("file", logFile),
	)
	
}

func getLogLevel() zapcore.Level {
	level:=config.GetLoggerConfig().Level
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

func Debugf(template string, args ...interface{}) {
	Sugar.Debugf(template, args...)
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Infof(template string, args ...interface{}) {
	Sugar.Infof(template, args...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Warnf(template string, args ...interface{}) {
	Sugar.Warnf(template, args...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Errorf(template string, args ...interface{}) {
	Sugar.Errorf(template, args...)
}

func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}

func Fatalf(template string, args ...interface{}) {
	Sugar.Fatalf(template, args...)
}