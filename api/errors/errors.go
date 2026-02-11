package errors

import (
	"mygo_bangforai/pkg/utils"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	//"fmt"
)

func NewError(code int,message string,op string) *ServerError {
	return &ServerError{
		Code:       code,
		Message:    message,
		Op:         op,
		Timestamp:  time.Now(),
		Context:    make(map[string]interface{}),
		Fields:     make(map[string]any),
		StackTrace: getStackTrace(),
	}
}

// WithContext 为错误添加上下文信息
func (e *ServerError) WithContext(key string, value interface{}) *ServerError {
	if e.Context == nil {
		e.Context = make(map[string]interface{})
	}
	e.Context[key] = value
	return e
}

// WithField 为错误添加字段信息
func (e *ServerError) WithField(key string, value interface{}) *ServerError {
	if e.Fields == nil {
		e.Fields = make(map[string]any)
	}
	e.Fields[key] = value
	return e
}

func WrapError(err error, code int, message string, op string) *ServerError {
	serverErr := &ServerError{
		Code:       code,
		Message:    message,
		Op:         op,
		Timestamp:  time.Now(),
		Original:   err,
		Context:    make(map[string]interface{}),
		Fields:     make(map[string]any),
		StackTrace: getStackTrace(),
	}
	handleError(*serverErr)
	return serverErr
}

func handleError(err ServerError) *ServerError {
	logger := utils.GetLogger()
	logger.Error(err.Op,zap.Int(err.Message,err.Code))

	
	return nil
}

// getStackTrace 获取当前的堆栈跟踪信息
func getStackTrace() string {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	// 过滤掉getStackTrace函数本身的调用
	stack := string(buf[:n])
	lines := strings.Split(stack, "\n")
	var filteredLines []string
	for _, line := range lines {
		if !strings.Contains(line, "getStackTrace") {
			filteredLines = append(filteredLines, line)
		}
	}
	return strings.Join(filteredLines, "\n")
}
