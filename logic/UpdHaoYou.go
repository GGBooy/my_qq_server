package logic

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdHaoYou(c *gin.Context) {
	hyId := c.DefaultPostForm("hyId", "0")
	hystatu, _ := strconv.Atoi(c.DefaultPostForm("hystatu", "0"))
	log.Printf("UpdHaoYou hyId:%s hystatu:%d\n", hyId, hystatu)

	qqhys := qqhy{}
	result := db.Find(&qqhys, "hy_id = ?", hyId)
	if result.Error != nil {
		log.Println("UpdHaoYou ERROR")
		log.Println(result.Error)
		c.JSON(http.StatusOK, gin.H{
			"result": 0,
		})
		return
	}

	qqhys.HyStatu = hystatu
	result = db.Save(&qqhys)
	if result.Error != nil {
		log.Println("UpdHaoYou ERROR")
		log.Println(result.Error)
		c.JSON(http.StatusOK, gin.H{
			"result": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": 1,
		})
	}
}
