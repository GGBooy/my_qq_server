package logic

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// "Android_Service/QQ/execute?zhanghao="+zhanghao+"&mima="+mima+""
func LoginHandler(c *gin.Context) {
	zhanghao, err := strconv.Atoi(c.DefaultQuery("zhanghao", "0"))
	if err != nil {
		log.Panicln("LOGIN_HANDLER: ", err)
	}
	mima := c.DefaultQuery("mima", "")

	type qquser qquserInDB
	qRes := qquser{}

	result := db.Where("qq_zhanghao = ?", zhanghao).Where("qq_mima = ?", mima).First(&qRes)
	if result.Error != nil {
		log.Println("login :", err)
		// send
		c.JSON(http.StatusOK, gin.H{
			"result":     0,
			"qqId":       qRes.QqId,
			"qqZhanghao": qRes.QqZhanghao,
			"qqMima":     qRes.QqMima,
			"qqTouxiang": qRes.QqTouxiang,
			"qqName":     qRes.QqName,
			"qqMark":     qRes.QqMark,
			"qqSex":      qRes.QqSex,
			"qqAddress":  qRes.QqAddress,
			"qqPhone":    qRes.QqPhone,
			"qqStatu":    qRes.QqStatu,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":     1,
		"qqId":       qRes.QqId,
		"qqZhanghao": qRes.QqZhanghao,
		"qqMima":     qRes.QqMima,
		"qqTouxiang": qRes.QqTouxiang,
		"qqName":     qRes.QqName,
		"qqMark":     qRes.QqMark,
		"qqSex":      qRes.QqSex,
		"qqAddress":  qRes.QqAddress,
		"qqPhone":    qRes.QqPhone,
		"qqStatu":    qRes.QqStatu,
	})
	log.Println("\n--Login", qRes.QqZhanghao, "has logged in\n")
}
