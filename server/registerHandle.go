package server

import (
	"my_qq_server/logic"

	"github.com/gin-gonic/gin"
)

func RegisterHandle() *gin.Engine {
	router := gin.Default()

	router.POST("Android_Service/QQ/zhuce", logic.ZhuceHandler)

	return router
	//
}
