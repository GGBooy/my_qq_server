package logic

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BecomeQqHyHandler(c *gin.Context) {
	type qqhyNoHyId struct {
		MyqqId       int    `json:"myqqId"`
		MyqqZhanghao string `json:"myqqZhanghao"`
		MyqqName     string `json:"myqqName"`
		MyqqTouxiang string `json:"myqqTouxiang"`
		MyqqMark     string `json:"myqqMark"`
		MyqqStatu    int    `json:"myqqStatu"`
		MyqqFengzu   int    `json:"myqqFengzu"`

		HyqqId       int    `json:"hyqqId"`
		HyqqZhanghao string `json:"hyqqZhanghao"`
		HyqqName     string `json:"hyqqName"`
		HyqqTouxiang string `json:"hyqqTouxiang"`
		HyqqMark     string `json:"hyqqMark"`
		HyqqStatu    int    `json:"hyqqStatu"`
		HyqqFengzu   int    `json:"hyqqFengzu"`
		HyStatu      int    `json:"hyStatu"`
	}
	type qqhy qqhyNoHyId
	myqqid, _ := strconv.Atoi(c.DefaultPostForm("qqhy.myqqId", "0"))
	myqqstatu, _ := strconv.Atoi(c.DefaultPostForm("qqhy.myqqStatu", "0"))
	myqqfengzu, _ := strconv.Atoi(c.DefaultPostForm("qqhy.myqqFengzu", "0"))

	hyqqid, _ := strconv.Atoi(c.DefaultPostForm("qqhy.hyqqId", "0"))
	hyqqstatu, _ := strconv.Atoi(c.DefaultPostForm("qqhy.hyqqStatu", "0"))
	hyqqfengzu, _ := strconv.Atoi(c.DefaultPostForm("qqhy.hyqqFengzu", "0"))

	hystatu, _ := strconv.Atoi(c.DefaultPostForm("qqhy.hyStatu", "0"))

	hy := qqhy{
		MyqqId:       myqqid,
		MyqqZhanghao: c.DefaultPostForm("qqhy.myqqZhanghao", ""),
		MyqqName:     c.DefaultPostForm("qqhy.myqqName", ""),
		MyqqTouxiang: c.DefaultPostForm("qqhy.myqqTouxiang", ""),
		MyqqMark:     c.DefaultPostForm("qqhy.myqqMark", ""),
		MyqqStatu:    myqqstatu,
		MyqqFengzu:   myqqfengzu,

		HyqqId:       hyqqid,
		HyqqZhanghao: c.DefaultPostForm("qqhy.hyqqZhanghao", ""),
		HyqqName:     c.DefaultPostForm("qqhy.hyqqName", ""),
		HyqqTouxiang: c.DefaultPostForm("qqhy.hyqqTouxiang", ""),
		HyqqMark:     c.DefaultPostForm("qqhy.hyqqMark", ""),
		HyqqStatu:    hyqqstatu,
		HyqqFengzu:   hyqqfengzu,
		HyStatu:      hystatu,
	}

	db.Create(&hy)
	log.Println("\n --becomeQqHY: ", hy.MyqqId, " -> ", hy.HyqqId, "\n")
}
