package controller

import (
	"github.com/gin-gonic/gin"
	"myweb/exception"
	"myweb/result"
	"net/http"
)

func HandlePanic(c *gin.Context) {
	err := recover()
	if err != nil {
		c.JSON(http.StatusOK, result.SysError("服务器异常"))
	}
}

func HandleError(c *gin.Context, err error) {
	if v, ok := err.(exception.BizError); ok {
		c.JSON(http.StatusOK, result.BizError(v.Code, v.Msg))
	} else {
		c.JSON(http.StatusOK, result.SysError("服务器异常"))
	}
}
