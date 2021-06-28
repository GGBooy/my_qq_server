package logic

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// "Android_Service/QQ/selhaoyou"
func SelHaoYouHandler(c *gin.Context) {
	qqId, err := strconv.Atoi(c.DefaultPostForm("qqId", ""))
	if err != nil {
		log.Println(err)

		return
	}

	hy := []qqhy{}
	result := db.Where("(myqq_id = ? OR hyqq_id = ?) AND hy_statu = ?", qqId, qqId, 1).Find(&hy)
	if result.Error != nil {
		log.Println("selHaoYou: ", err)
		content := gin.H{
			"result":     0,
			"applycount": 0,
			"qqhy":       nil,
		}
		c.JSON(http.StatusOK, content)
	}

	content := gin.H{
		"result":     1,
		"applycount": 0, // TODO -- WHAT THE HELL IS THIS?
		"qqhy":       hy,
	}
	c.JSON(http.StatusOK, content)

	log.Println("\n--selHaoYou", qqId, "\n")
}
