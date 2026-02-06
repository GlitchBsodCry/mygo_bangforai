package errors

import (
	"fmt"
	"runtime"
)

// Error 是一个基础错误类型
type Error0 struct {
	Message    string
	Code       int
	Original   error
	StackTrace string
}

// Error 实现 error 接口
func (e *Error0) Error() string {
	if e.Original != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Original)
	}
	return e.Message
}

// Unwrap 返回原始错误，用于 errors.Is 和 errors.As 函数
func (e *Error0) Unwrap() error {
	return e.Original
}

// GetStackTrace 获取错误堆栈
func (e *Error0) GetStackTrace() string {
	return e.StackTrace
}


// Wrap 包装一个错误，可选择性添加错误码
func Wrap(err error, message string, code ...int) error {
	if err == nil {
		return nil
	}
	errorCode := 0
	if len(code) > 0 {
		errorCode = code[0]
	}
	return &Error0{
		Message:    message,
		Code:       errorCode,
		Original:   err,
		StackTrace: getStackTrace(),
	}
}

// Is 判断错误是否为指定类型
func Is(err, target error) bool {
	if err == target {
		return true
	}
	if e, ok := err.(*Error0); ok {
		return Is(e.Original, target)
	}
	return false
}

// As 将错误转换为指定类型
func As(err error, target interface{}) bool {
	if err == target {
		return true
	}
	if e, ok := err.(*Error0); ok {
		return As(e.Original, target)
	}
	return false
}

// getStackTrace 获取当前堆栈信息
func getStackTrace() string {
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}
