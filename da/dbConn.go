package da

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbConn *gorm.DB
var dbmx sync.Mutex

func CloseDB() {
	err := dbConn.Close()
	if err != nil {
		log.Println("Error in Closing DB", err)
	}
}

//参数含义:数据库用户名、密码、主机ip、连接的数据库、端口号
func createDBConn(User, Password, Host, Db string, Port int) *gorm.DB {
	// TODO charset
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", User, Password, Host, Port, Db)
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Println("\n\n\n\ncreateDBConn", err)
		return nil
	}
	// db.SingularTable(true) //如果使用gorm来帮忙创建表时，这里填写false的话gorm会给表添加s后缀，填写true则不会
	// db.LogMode(true)       //打印sql语句

	//开启连接池
	db.DB().SetMaxIdleConns(20)                   //最大空闲连接
	db.DB().SetMaxOpenConns(100)                  //最大连接数
	db.DB().SetConnMaxLifetime(600 * time.Second) //最大生存时间(s)

	return db
}

// GetDB 开放给外部获得db连接
func GetDB() *gorm.DB {
	dbmx.Lock()
	defer dbmx.Unlock()

	if dbConn == nil {
		dbConn = createDBConn("root", "....", "192.168.3.3", "androidqqdata", 3306)
	}
	return dbConn
}

// func GetDb() *gorm.DB {
// 	var conn *gorm.DB = nil
// 	for {
// 		conn = dbConn("root", "....", "192.168.3.3", "androidqqdata", 3306)
// 		if conn != nil {
// 			break
// 		}
// 		fmt.Println("本次未获取到mysql连接")
// 	}
// 	return conn
// }
