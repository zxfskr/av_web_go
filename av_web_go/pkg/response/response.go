package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var statusCode webStatusCode

type webStatusCode struct {
	OK        string
	ERROR     string
	EXCEPTION string
	LOGINFAIL string
}

func init() {
	statusCode.OK = "000000"
	statusCode.ERROR = "900000"
	statusCode.EXCEPTION = "900001"
	statusCode.LOGINFAIL = "200001"
}

// Successful 成功请求
func Successful(c *gin.Context, msg string, data gin.H) {
	if msg == "" {
		msg = "操作成功"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    statusCode.OK,
		"message": msg,
		"data":    data,
	})
}

// Fail 请求失败
func Fail(c *gin.Context, msg string) {
	if msg == "" {
		msg = "操作失败"
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    statusCode.ERROR,
		"message": msg,
		"data":    nil,
	})
}

// Exception 系统异常
func Exception(c *gin.Context, msg string) {
	if msg == "" {
		msg = "系统异常"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    statusCode.EXCEPTION,
		"message": msg,
		"data":    nil,
	})
}

// LoginFail 登录异常
func LoginFail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    statusCode.LOGINFAIL,
		"message": "登录失败",
		"data":    nil,
	})
}
