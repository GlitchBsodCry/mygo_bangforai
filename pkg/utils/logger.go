package utils

import (
	"go.uber.org/zap"
)

// LoggerInterface 定义日志接口，用于解决循环依赖问题
type LoggerInterface interface {
	Error(msg string, fields ...zap.Field)
	Errorf(template string, args ...interface{})
}

// LoggerInstance 全局日志实例
var LoggerInstance LoggerInterface

// SetLogger 设置日志实例
func SetLogger(logger LoggerInterface) {
	LoggerInstance = logger
}

// GetLogger 获取日志实例
func GetLogger() LoggerInterface {
	return LoggerInstance
}