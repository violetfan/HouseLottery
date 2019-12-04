package model

import "time"

// 购房用户数据测试
var UserStore = map[string]User{
	"Lena": {
		ID:           201600100,  //审核编号
		Name:         "Lena",
		HeadImg:      "",
		Sex:          "女",
		Phone:        "13845697412",
		Password:     "789654",
		Right:        0,
		RegisterTime: time.Unix(1464134400,0),  //2016/05/25
		Status:       0,
		HouseType:    "花园房",  //意向房源类型
		PurposeHouse: "3B-502-1",  //意向房号
		IdentityType: "居民身份证",   //证件类型
		IdentityNum:  "610124198905165226",  //证件号码
		CheckStatus:  0,  //审核状态
	},
	"Finn": {
		ID:           201700123,
		Name:         "Finn",
		HeadImg:      "",
		Sex:          "男",
		Phone:        "18223654785",
		Password:     "539654",
		Right:        0,
		RegisterTime: time.Unix(1485100800,0),   //2017/01/23
		Status:       1,
		HouseType:    "普通小区房",
		PurposeHouse: "1A-403-2",
		IdentityType: "居民身份证",
		IdentityNum:  "610124197107074125",
		CheckStatus:  0,
	},
	"Jonas": {
		ID:           201801250,
		Name:         "Jonas",
		HeadImg:      "",
		Sex:          "男",
		Phone:        "15674128521",
		Password:     "369745",
		Right:        0,
		RegisterTime: time.Unix(1543939200,0),  //2018/12/05
		Status:       1,
		HouseType:    "老居民房",
		PurposeHouse: "1B-304-5",
		IdentityType: "居民身份证",
		IdentityNum:  "610526197305217534",
		CheckStatus:  1,
	},
	"Emma": {
		ID:           201400342,
		Name:         "Noah",
		HeadImg:      "",
		Sex:          "女",
		Phone:        "18245612354",
		Password:     "753951",
		Right:        1,
		RegisterTime: time.Unix(1395590400,0),  //2014/03/24
		Status:       1,
		HouseType:    "集资房",
		PurposeHouse: "5A-607-2",
		IdentityType: "居民身份证",
		IdentityNum:  "610124198502124562",
		CheckStatus:  0,
	},
	"Noah": {
		ID:           201535617,
		Name:         "Noah",
		HeadImg:      "",
		Sex:          "男",
		Phone:        "13896542357",
		Password:     "852364",
		Right:        1,
		RegisterTime: time.Unix(1433001600,0),   //2015/5/31
		Status:       0,
		HouseType:    "经适房",
		PurposeHouse: "8B-201-3",
		IdentityType: "居民身份证",
		IdentityNum:  "610526197105272566",
		CheckStatus:  1,
	},
}
