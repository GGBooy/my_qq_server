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

	router.POST("Android_Service/QQ/selhaoyouOfmy", logic.SelhaoyouOfmy)

	router.POST("Android_Service/QQ/seluserlist", logic.Seluserlist)

	router.POST("Android_Service/QQ/selqquser", logic.Selqquser)

	router.POST("Android_Service/QQ/becomeQqhy", logic.BecomeQqHyHandler)

	router.POST("Android_Service/QQ/UpdHaoYou", logic.UpdHaoYou)

	router.POST("Android_Service/QQ/sendMessage", logic.SendMessage)

	router.POST("Android_Service/QQ/receiveMessage", logic.ReceiveMessage)
	return router
	//
}
