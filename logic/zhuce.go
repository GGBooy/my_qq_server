package logic

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// TODO
// AUTO REGISTER(OPEN-CLOSE PRINCIPLE)

type qquser struct {
	qqTouxiang string
	qqPhone    string
	qqName     string
	qqMima     string
	qqSex      string
	qqAddress  string
	qqMark     string
	qqStatu    string
	qqZhanghao string
}

func ZhuceHandler(c *gin.Context) {
	qu := qquser{
		qqTouxiang: c.DefaultPostForm("qquser.qqTouxiang", ""),
		qqPhone:    c.DefaultPostForm("qquser.qqPhone", ""),
		qqName:     c.DefaultPostForm("qquser.qqName", ""),
		qqMima:     c.DefaultPostForm("qquser.qqMima", ""),
		qqSex:      c.DefaultPostForm("qquser.qqSex", ""),
		qqAddress:  c.DefaultPostForm("qquser.qqAddress", ""),
		qqMark:     c.DefaultPostForm("qquser.qqMark", ""),
		qqStatu:    c.DefaultPostForm("qquser.qqStatu", ""),
		qqZhanghao: c.DefaultPostForm("qquser.qqZhanghao", ""),
	}

	fmt.Println(qu)
}
