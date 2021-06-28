package logic

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// "Android_Service/QQ/zhuce"
func ZhuceHandler(c *gin.Context) {
	type qquser qquserSignUp
	qu := qquser{
		QqTouxiang: c.DefaultPostForm("qquser.qqTouxiang", ""),
		QqPhone:    c.DefaultPostForm("qquser.qqPhone", ""),
		QqName:     c.DefaultPostForm("qquser.qqName", ""),
		QqMima:     c.DefaultPostForm("qquser.qqMima", ""),
		QqSex:      c.DefaultPostForm("qquser.qqSex", ""),
		QqAddress:  c.DefaultPostForm("qquser.qqAddress", ""),
		QqMark:     c.DefaultPostForm("qquser.qqMark", ""),
		QqStatu:    0,
		QqZhanghao: c.DefaultPostForm("qquser.qqZhanghao", ""),
	}

	status, _ := strconv.Atoi(c.DefaultPostForm("qquser.qqStatu", ""))
	qu.QqStatu = status

	db.Create(&qu)

	log.Println("\n--ZHUCE: ", qu, "\n")
}
