package model

import "time"

//员工数据测试
var StaffStore = map[string]Staff{
	"Admin": {
		ID:          001,
		Name:        "Admin",
		HeadImg:     "——",
		Sex:         "男",
		Phone:       "11111",
		Password:    "123456",
		Right:       7,
		Status: 0,
		JoinTime:    time.Unix(1462838400,0),  //员工加入时间   2016/5/10
	},
	"Leonie": {
		ID:          002,
		Name:        "Leonie",
		HeadImg:     "__",
		Sex:         "女",
		Phone:       "13812341234",
		Password:    "147852",
		Right:       1,
		Status: 1,
		JoinTime:    time.Unix(1476057600,0),   //2016/10/10
	},
	"David": {
		ID:          003,
		Name:        "David",
		HeadImg:     "__",
		Sex:         "男",
		Phone:       "18236981254",
		Password:    "951357",
		Right:       4,
		Status: 1,
		JoinTime:    time.Unix(1427068800,0),   //2015/3/23
	},
	"Tim": {
		ID:          004,
		Name:        "Tim",
		HeadImg:     "__",
		Sex:         "男",
		Phone:       "18275369852",
		Password:    "823975",
		Right:       4,
		Status: 0,
		JoinTime:    time.Unix(1409184000,0),  //2014/8/28
	},
	"Mia": {
		ID:          005,
		Name:        "Mia",
		HeadImg:     "__",
		Sex:         "女",
		Phone:       "15675894562",
		Password:    "856742",
		Right:       2,
		Status: 0,
		JoinTime:    time.Unix(1464134400,0),  //2016/5/25
	},
}

