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
	// result - for reply purpose 1 for ok
	Result int
}

// qqhy
type qqhy struct {
	hy_id         int
	myqq_id       int
	myqq_zhanghao string
	myqq_name     string
	myqq_touxiang string
	myqq_mark     string
	myqq_statu    int
	myqq_fengzu   int
	hyqq_id       int
	hyqq_zhanghao string
	hyqq_name     string
	hyqq_touxiang string
	hyqq_mark     string
	hyqq_statu    int
	hyqq_fengzu   int
	hy_statu      int
}
