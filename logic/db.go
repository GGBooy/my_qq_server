package logic

import (
	"my_qq_server/da"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db = da.GetDB()
}

func CloseDB() {
	da.CloseDB()
}
