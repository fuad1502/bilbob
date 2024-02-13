package errors

import "github.com/gin-gonic/gin"

type GinError struct {
	err         error
	handlerName string
	funcName    string
}

func New(err error, c *gin.Context, funcName string) *GinError {
	return &GinError{
		err:         err,
		handlerName: c.HandlerName(),
		funcName:    funcName,
	}
}

func (e *GinError) Error() string {
	return e.handlerName + ": " + e.funcName + ": " + e.err.Error()
}
