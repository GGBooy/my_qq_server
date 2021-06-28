package server

import (
	"my_qq_server/logic"

	"github.com/gin-gonic/gin"
)

func RegisterHandle() *gin.Engine {
	router := gin.Default()

	router.POST("Android_Service/QQ/zhuce", logic.ZhuceHandler)

	router.GET("Android_Service/QQ/execute", logic.LoginHandler)

	router.POST("Android_Service/QQ/selhaoyou", logic.SelHaoYouHandler)

	return router
	//
}
