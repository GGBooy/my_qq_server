package logic

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func SelhaoyouOfmy(c *gin.Context) {
	qqId, err := strconv.Atoi(c.DefaultQuery("qqId", "0"))
	if err != nil {
		log.Panicln("LOGIN_HANDLER: ", err)
	}
	log.Println("selhaoyouOfmy: qqId %d", qqId)

	// selQqhyOfmy
	listcount := []qqhy{}
	result := db.Where(&qqhy{hy_id: qqId, hy_statu: 0}).Find(&listcount)
	if result.Error != nil {
		resultfail(c, result.Error)
		return
	}

	// selQqhyOfApplyLit
	list := []qqhy{}
	result = db.Find(&list, "hyqq_id = ? AND hy_statu BETWEEN ? AND ? "+
		"OR myqq_id = ? AND hy_statu BETWEEN ? AND ?", qqId, 0, 2, qqId, 0, 2)
	if result.Error != nil {
		resultfail(c, result.Error)
		return
	}

	if len(list) != 0 {
		hysqlist := []qqhy{}
		for i := 0; i < len(list); i++ {
			qqhyItem := qqhy{}
			result = db.Find(&qqhyItem, "hy_id = ?", list[i].hy_id)
			if result.Error != nil {
				resultfail(c, result.Error)
				return
			}
			hysqlist = append(hysqlist, qqhyItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"result":     1,
			"qqhy":       hysqlist,
			"applycount": len(listcount),
		})
	} else {
		resultfail(c, nil)
	}
}

func resultfail(c *gin.Context, err error) {
	log.Println("SelhaoyouOfmy ERROR")
	log.Println(err)
	c.JSON(http.StatusOK, gin.H{
		"result": 0,
	})
}
