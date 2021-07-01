package logic

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReceiveMessage(c *gin.Context) {
	qqId, err := strconv.Atoi(c.DefaultPostForm("qqId", "0"))
	if err != nil {
		log.Println(err)

		return
	}
	hyId, err := strconv.Atoi(c.DefaultPostForm("hyId", "0"))
	if err != nil {
		log.Println(err)

		return
	}

	msgs := []qqmessage{}
	result := db.Where("(qq_id = ? AND m_jsid = ?) OR (qq_id = ? AND m_jsid = ?)", qqId, hyId, hyId, qqId).Find(&msgs)
	if result.Error != nil {
		log.Println("receiveMsg: ", err)
		content := gin.H{
			"result": 0,
			"msg":    nil,
		}
		c.JSON(http.StatusOK, content)
	}

	content := gin.H{
		"result": 1,
		"msg":    msgs,
	}
	c.JSON(http.StatusOK, content)

	log.Println("\n--ReceiveMsg", qqId, "->", hyId, "\n")
}
