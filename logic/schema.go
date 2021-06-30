package logic

import "time"

// the qquer used in sign up
type qquserSignUp struct {
	QqTouxiang string `json:"qqTouxiang"`
	QqPhone    string `json:"qqPhone"`
	QqName     string `json:"qqName"`
	QqMima     string `json:"qqMima"`
	QqSex      string `json:"qqSex"`
	QqAddress  string `json:"qqAddress"`
	QqMark     string `json:"qqMark"`
	QqStatu    int    `json:"qqStatu"`
	QqZhanghao string `json:"qqZhanghao"`
}

// the qquser in the db
type qquserInDB struct {
	// qquserData
	QqTouxiang string `json:"qqTouxiang"`
	QqPhone    string `json:"qqPhone"`
	QqName     string `json:"qqName"`
	QqMima     string `json:"qqMima"`
	QqSex      string `json:"qqSex"`
	QqAddress  string `json:"qqAddress"`
	QqMark     string `json:"qqMark"`
	QqStatu    int    `json:"qqStatu"`
	QqZhanghao string `json:"qqZhanghao"`
	// id - the primary key
	QqId int `json:"qqId"`
}

// the qquser in the frontend has two more fields
type qquserLoginReply struct {
	// qquserData
	QqTouxiang string `json:"qqTouxiang"`
	QqPhone    string `json:"qqPhone"`
	QqName     string `json:"qqName"`
	QqMima     string `json:"qqMima"`
	QqSex      string `json:"qqSex"`
	QqAddress  string `json:"qqAddress"`
	QqMark     string `json:"qqMark"`
	QqStatu    int    `json:"qqStatu"`
	QqZhanghao string `json:"qqZhanghao"`
	// id - the primary key
	QqId int `json:"qqId"`
	// result - for reply purpose, 1 for ok
	Result int `json:"result"`
}

// qqhy
type qqhy struct {
	HyId int `json:"hyId" gorm:"primary_key"`

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

type qqmessage struct {
	MId        int       `json:"MId"`
	QqId       int       `json:"qqId"`
	QqZhanghao string    `json:"qqZhanghao"`
	QqName     string    `json:"qqName"`
	QqTouxiang string    `json:"qqTouxiang"`
	MMessage   string    `json:"MMessage"`
	MDate      time.Time `json:"MDate"`
	MJsid      int       `json:"MJsid"`
	MZhanghao  string    `json:"MZhanghao"`
	MName      string    `json:"MName"`
	MTouxiang  string    `json:"MTouxiang"`
	MStatu     int       `json:"MStatu"`
	result     int       `json:"result"`
}

type qqshowmessage struct {
	SmId         int       `json:"smId"`
	QqId         int       `json:"qqId"`
	QqZhanghao   string    `json:"qqZhanghao"`
	QqName       string    `json:"qqName"`
	QqTouxiang   string    `json:"qqTouxiang"`
	HyqqId       int       `json:"hyqqId"`
	HyqqZhanghao string    `json:"hyqqZhanghao"`
	HyqqName     string    `json:"hyqqName"`
	HyqqTouxiang string    `json:"hyqqTouxiang"`
	SmContent    string    `json:"smContent"`
	SmDate       time.Time `json:"smDate"`
}
