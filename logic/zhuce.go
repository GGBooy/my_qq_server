package logic

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ZhuceHandler(c *gin.Context) {
	qu := qquser{
		QqTouxiang: c.DefaultPostForm("qquser.qqTouxiang", ""),
		QqPhone:    c.DefaultPostForm("qquser.qqPhone", ""),
		QqName:     c.DefaultPostForm("qquser.qqName", ""),
		QqMima:     c.DefaultPostForm("qquser.qqMima", ""),
		QqSex:      c.DefaultPostForm("qquser.qqSex", ""),
		QqAddress:  c.DefaultPostForm("qquser.qqAddress", ""),
		QqMark:     c.DefaultPostForm("qquser.qqMark", ""),
		QqStatu:    c.DefaultPostForm("qquser.qqStatu", ""),
		QqZhanghao: c.DefaultPostForm("qquser.qqZhanghao", ""),
	}
	if db == nil {
		fmt.Println("HELL NO!!")
	}

	// temp := qquser{}
	// db.First(&temp)
	// fmt.Println(temp)
	// db1 := da.GetDB()
	// if db1 == nil {
	// 	fmt.Println("HELL NO NONONOONON!!")
	// }
	db.Create(&qu)

	fmt.Println(qu)
}
