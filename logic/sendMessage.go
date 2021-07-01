package logic

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context) {
	qqid, _ := strconv.Atoi(c.DefaultPostForm("msg.qqId", "0"))
	mjsid, _ := strconv.Atoi(c.PostForm("msg.MJsid"))
	status, _ := strconv.Atoi(c.PostForm("msg.MStatu"))
	msg := qqmessage{
		QqId:       qqid,
		QqZhanghao: c.PostForm("msg.qqZhanghao"),
		QqName:     c.PostForm("msg.qqName"),
		QqTouxiang: c.PostForm("msg.qqTouxiang"),
		MMessage:   c.PostForm("msg.MMessage"),
		MJsid:      mjsid,
		MZhanghao:  c.PostForm("msg.MZhanghao"),
		MName:      c.PostForm("msg.MName"),
		MTouxiang:  c.PostForm("msg.MTouxiang"),
		MStatu:     status,
	}
	msg.MDate = time.Now().UnixNano() / 1e6
	result := db.Create(&msg)
	if result.Error != nil {
		log.Println("sendMessage1 ERROR")
		log.Println(result.Error)
		c.JSON(http.StatusOK, gin.H{
			"result": 0,
		})
		return
	}

	showmsg := []qqshowmessage{}
	db.Find(&showmsg, "(qq_id = ? and hy_id = ?) or (qq_id = ? and hy_id = ?)",
		msg.QqId, msg.MJsid, msg.MJsid, msg.QqId)
	if showmsg != nil && len(showmsg) != 0 {
		qmsg := qqshowmessage{}
		db.Find(&qmsg, "sm_id = ?", showmsg[0].SmId)
		qmsg.SmContent = msg.MMessage
		db.Save(&qmsg)
	} else {
		aqsmsg := qqshowmessage{
			QqId:       msg.QqId,
			QqZhanghao: msg.QqZhanghao,
			QqName:     msg.QqName,
			QqTouxiang: msg.QqTouxiang,
			HyId:       msg.MJsid,
			HyZhanghao: msg.MZhanghao,
			HyName:     msg.MName,
			HyTouxiang: msg.MTouxiang,
			SmContent:  msg.MMessage,
			SmDate:     time.Now().UnixNano() / 1e6,
		}
		result = db.Create(&aqsmsg)
		if result.Error != nil {
			log.Println("sendMessage2 ERROR")
			log.Println(result.Error)
			c.JSON(http.StatusOK, gin.H{
				"result": 0,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"result": 1,
	})

}
