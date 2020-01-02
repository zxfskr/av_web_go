package user

import (
	"av_process/av_web_go/model"
	"av_process/av_web_go/pkg/captcha"
	"av_process/av_web_go/pkg/response"
	"av_process/common"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

// InitUserView 注册user视图
func InitUserView(g *gin.RouterGroup) {
	g.GET("/test", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Options(sessions.Options{
			MaxAge: 10,
		})
		s := session.Get("login_user")
		fmt.Println(s)
		session.Set("login_user", "haha")
		session.Save()
		response.Successful(c, "", nil)
	})
	logger = common.Logger
	g.GET("/logout", logout)
	g.GET("/list", getAllUsersByPage)
}

func getAllUsersByPage(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")

	numLimit, err := strconv.Atoi(limit)
	if err != nil {
		logger.Error(err)
		response.Fail(c, "参数limit异常")
		return
	}
	numOffset, err := strconv.Atoi(offset)
	if err != nil {
		logger.Error(err)
		response.Fail(c, "参数offset异常")
		return
	}
	userModel := model.UserModel{}
	users, err := userModel.GetAllByPage(numLimit, numOffset)
	if err != nil {
		logger.Error(err)
		response.Fail(c, "查找用户失败")
		return
	}

	response.Successful(c, "", gin.H{"rows": users})
}

// GenCaptcha  生成验证码
func GenCaptcha(c *gin.Context) {
	img, text, err := captcha.CreateCaptcha()
	if err != nil {
		logger.Error("验证码图片生成失败， err: ", err)
		response.Exception(c, "")
	}

	session := sessions.Default(c)
	session.Set("captcha", strings.ToLower(text))
	session.Save()
	logger.Debug("captcha: ", strings.ToLower(text))

	contentType := "image/jpg"
	c.Data(http.StatusOK, contentType, img)
}

// Login 登录
func Login(c *gin.Context) {
	userModel := model.UserModel{}
	username := c.Query("username")
	password := c.Query("password")
	captcha := c.Query("captcha")

	user, err := userModel.GetOneByName(username)
	if err != nil {
		logger.Error(err)
		response.Fail(c, "查找用户失败")
		return
	}

	pdMD5 := common.GenMD5(password)
	if pdMD5 != user.Password {
		response.Fail(c, "密码错误")
		return
	}

	session := sessions.Default(c)
	// session.Options(sessions.Options{
	// 	MaxAge: 3600,
	// })

	fmt.Println(strings.ToLower(captcha), session.Get("captcha"))
	if strings.ToLower(captcha) != session.Get("captcha") {
		response.Fail(c, "验证码错误")
		return
	}

	session.Set("userID", user.UserID)
	session.Save()
	logger.Info("用户", username, " 登录")

	response.Successful(c, "登录成功", gin.H{"userType": user.UserType})
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	response.Successful(c, "退出登录", nil)
}
