package logic

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Selqquser(c *gin.Context) {
	qqId := c.DefaultPostForm("qqId", "")
	log.Println("selqquser: qqId", qqId)

	type qquser qquserInDB

	user := qquser{}
	result := db.First(&user, "qq_id = ?", qqId)
	if result.Error != nil {
		log.Println("selqquser ERROR")
		log.Println(result.Error)
		c.JSON(http.StatusOK, gin.H{
			"result": 0,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result":     1,
		"qqId":       user.QqId,
		"qqZhanghao": user.QqZhanghao,
		"qqMima":     user.QqMima,
		"qqTouxiang": user.QqTouxiang,
		"qqName":     user.QqName,
		"qqMark":     user.QqMark,
		"qqSex":      user.QqSex,
		"qqAddress":  user.QqAddress,
		"qqPhone":    user.QqPhone,
		"qqStatu":    user.QqStatu,
	})
}
