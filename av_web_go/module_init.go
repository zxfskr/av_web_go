package main

import (
	"av_process/av_web_go/pkg/db"
	"av_process/av_web_go/view"

	"github.com/gin-gonic/gin"
)

// createRouter 创建route
func createRouter() (router *gin.Engine) {
	router = gin.New()
	// router := gin.Default()
	view.ModuleInit(router)
	db.ModuleInit()
	return
}
