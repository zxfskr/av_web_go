package view

import (
	"av_process/av_web_go/model"
	"av_process/av_web_go/pkg/response"
	"av_process/av_web_go/view/user"
	"av_process/common"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry
var router *gin.Engine

// ModuleInit 初始化模块
func ModuleInit(engine *gin.Engine) *gin.Engine {
	logger = common.Logger
	router = engine
	router.Use(gin.Recovery(), loggerToFile())
	store := cookie.NewStore([]byte("secret_av_web+"))
	router.Use(sessions.Sessions("av_web_session", store))
	// router.POST("/auth-web/sys/login", login)

	router.GET("/captcha", user.GenCaptcha)
	router.POST("/login", user.Login)

	view := router.Group("/")
	view.Use(authMiddleware())
	user.InitUserView(view)
	// 404 Handler.
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	return router
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		session := sessions.Default(c)

		userID := session.Get("userID")
		if userID == nil {
			response.LoginFail(c)
			c.Abort()
			return
		}
		userModel := model.UserModel{}
		_, err := userModel.GetOneByID(userID.(int))
		if err != nil {
			response.LoginFail(c)
			c.Abort()
			return
		}

		c.Next()
	}
}

// LoggerToFile 日志记录到文件
func loggerToFile() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqURI := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqURI,
		)
	}
}
