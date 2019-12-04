package model

import "time"

var StaffStore = map[string]Staff{
	"Admin": {
		ID:          0,
		Name:        "Admin",
		HeadImg:     "",
		Sex:         "男",
		Phone:       "11111",
		Password:    "123456",
		Right:       7,
		StaffStatus: 0,
		JoinTime:    time.Time{},
	},
	"Lirui": {
		ID:          0,
		Name:        "Lirui",
		HeadImg:     "__",
		Sex:         "女",
		Phone:       "33333",
		Password:    "123321",
		Right:       7,
		StaffStatus: 1,
		JoinTime:    time.Time{},
	},
}

var UserStore = map[string]User{
	"wangchengy": {
		ID:           0,
		Name:         "",
		HeadImg:      "",
		Sex:          "",
		Phone:        "",
		Password:     "",
		Right:        0,
		RegisterTime: time.Time{},
		Status:       0,
		HouseType:    "",
		PurposeHouse: "",
		IdentityType: "",
		IdentityNum:  "",
		CheckStatus:  0,
	},
}