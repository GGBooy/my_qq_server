package logic

// the qquer used in sign up
type qquserSignUp struct {
	QqTouxiang string
	QqPhone    string
	QqName     string
	QqMima     string
	QqSex      string
	QqAddress  string
	QqMark     string
	QqStatu    int
	QqZhanghao string
}

// the qquser in the db
type qquserInDB struct {
	// qquserData
	QqTouxiang string
	QqPhone    string
	QqName     string
	QqMima     string
	QqSex      string
	QqAddress  string
	QqMark     string
	QqStatu    int64
	QqZhanghao string
	// id - the primary key
	QqId int64
}

// the qquser in the frontend has two more fields
type qquserLoginReply struct {
	// qquserData
	QqTouxiang string
	QqPhone    string
	QqName     string
	QqMima     string
	QqSex      string
	QqAddress  string
	QqMark     string
	QqStatu    int64
	QqZhanghao string
	// id - the primary key
	QqId int64
	// result - for reply purpose, 1 for ok
	Result int
}
