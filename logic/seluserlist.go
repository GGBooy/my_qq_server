package logic

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Seluserlist(c *gin.Context) {
	xinxi := c.DefaultPostForm("xinxi", "")
	log.Println("selhaoyouOfmy: xinxi: ", xinxi)

	type qquser qquserInDB

	list := []qquser{}
	result := db.Find(&list, "qq_zhanghao like ? OR qq_name like ?", "%"+xinxi+"%", "%"+xinxi+"%")
	if result.Error != nil {
		resultfail(c, result.Error)
		return
	}

	userlists := []qquser{}
	if len(list) != 0 {
		for i := 0; i < len(list); i++ {
			user := qquser{}
			result := db.Find(&user, "qq_id = ?", list[i].QqId)
			if result.Error != nil {
				resultfail(c, result.Error)
				return
			}
			userlists = append(userlists, user)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": 1,
			"qquser": userlists,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": 0,
		})
	}
}
