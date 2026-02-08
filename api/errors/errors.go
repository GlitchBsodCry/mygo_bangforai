package errors

import (
	//"fmt"
	"mygo_bangforai/pkg/logger"
	"time"
)

func NewError(code int, message string, op string, original error, context map[string]interface{}, fields map[string]any) *ServerError {
	return &ServerError{
		Code:       code,
		Message:    message,
		Op:         op,
		Timestamp:  time.Now(),
		Original:   original,
		Context:    context,
		Fields:     fields,
	}
}

func WrapError(err error, code int, message string, op string, context map[string]interface{}, fields map[string]any) *ServerError {
	return &ServerError{
		Code:       code,
		Message:    message,
		Op:         op,
		Timestamp:  time.Now(),
		Original:   err,
		Context:    context,
		Fields:     fields,
	}
}

func HandleError(err ServerError) *ServerError {
	logger.Error(err.Message, logger.Zap.Any("error", err))
	return nil
}