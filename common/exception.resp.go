package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiException struct {
	// 0 正常
	// nil 未知状态
	ErrorCode *int   `json:"error_code"`
	HttpCode  int    `json:"http_code"`
	Message   string `json:"message"`
}

func NewApiException(errCode, httpCode int, msg string) *ApiException {
	return &ApiException{
		ErrorCode: &errCode,
		HttpCode:  httpCode,
		Message:   msg,
	}
}

func (e *ApiException) SetHttpCode(code int) *ApiException {
	e.HttpCode = code
	return e
}

// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.

//	type error interface {
//		Error() string
//	}
func (e *ApiException) Error() string {
	return e.Message
}

func RespFail(c *gin.Context, err error) {
	respMsg, ok := err.(*ApiException)
	if ok {
		c.JSON(respMsg.HttpCode, NewApiException(*respMsg.ErrorCode, respMsg.HttpCode, respMsg.Message))
		return
	}
	c.JSON(http.StatusInternalServerError, NewApiException(-1, http.StatusInternalServerError, respMsg.Error()))
}
