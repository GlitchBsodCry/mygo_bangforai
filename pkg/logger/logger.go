package logger

import (
	"os"
	"path/filepath"

	"mygo_bangforai/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// Zap 封装zap的字段函数
type ZapType struct{}

var Zap ZapType

// String 创建字符串类型的字段
func (z ZapType) String(key, value string) zap.Field {
	return zap.String(key, value)
}

// Int 创建整数类型的字段
func (z ZapType) Int(key string, value int) zap.Field {
	return zap.Int(key, value)
}

// Bool 创建布尔类型的字段
func (z ZapType) Bool(key string, value bool) zap.Field {
	return zap.Bool(key, value)
}

// Error 创建错误类型的字段
func (z ZapType) Error(err error) zap.Field {
	return zap.Error(err)
}

func InitLogger() {
	// 从配置中获取日志配置
	loggerConfig := config.GetLoggerConfig()
	
	// 设置日志级别
	logLevel := getLogLevel(loggerConfig.Level)
	
	// 设置编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	
	var encoder zapcore.Encoder
	if loggerConfig.Format == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}
	
	// 设置输出
	var writer zapcore.WriteSyncer
	if loggerConfig.Output == "file" {
		// 确保日志目录存在
		logDir := filepath.Dir(loggerConfig.File)
		os.MkdirAll(logDir, 0755)
		
		// 打开日志文件
		file, err := os.OpenFile(loggerConfig.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			// 如果文件打开失败，使用标准输出
			writer = zapcore.AddSync(os.Stdout)
			Info("日志文件打开失败，使用标准输出", Zap.Error(err))
		} else {
			writer = zapcore.AddSync(file)
		}
	} else {
		writer = zapcore.AddSync(os.Stdout)
	}
	
	// 创建核心
	core := zapcore.NewCore(
		encoder,
		writer,
		logLevel,
	)

	// 创建日志器
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	
	Info("日志初始化完成", 
		Zap.String("level", loggerConfig.Level),
		Zap.String("format", loggerConfig.Format),
		Zap.String("output", loggerConfig.Output),
		Zap.String("file", loggerConfig.File),
	)
}

// getLogLevel 根据配置的日志级别字符串返回对应的zapcore.Level
func getLogLevel(level string) zapcore.Level {
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

func GetLogger() *zap.Logger {
	return Logger
}

func Info(msg string, fields ...zap.Field) {// 信息级别
	Logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {// 错误级别
	Logger.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {// 调试级别
	Logger.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {// 警告级别
	Logger.Warn(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {// 致命错误级别
	Logger.Fatal(msg, fields...)
}