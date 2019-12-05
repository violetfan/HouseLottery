package model

import "time"

//员工 type=0
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

//房屋信息
var HouseStore = map[string]House{
	"风栖云筑": {
		BuildID:          "201912365",              //楼盘号
		BuildName:        "风栖云筑",                   //楼盘名
		ChooseStime:      time.Unix(1575417600, 0), //选房开始时间   2019/12/04 08:00:00
		ChooseEndtime:    time.Unix(1575540000, 0), //选房结束时间   2019/12/05 18:00:00
		Enterprise:       "西安耀藏置业有限公司",             //开发企业
		Introduce:        "有凤来仪非梧不栖，中轴凤栖，欢迎品鉴",     //介绍
		Hotline:          "4006708280",             //咨询电话
		PresellNumber:    "41407",                  //预售证号
		IntentStime:      time.Unix(1572912000, 0), //意向登记开始时间   //2019/11/05 8:00:00
		IntentEndtime:    time.Unix(1573207200, 0), //意向登记结束时间  //2019/11/08 18:00:00
		ReceptionStime:   time.Unix(1575936000, 0), //现场接受资料开始时间   //2019/12/10 8:00:00
		ReceptionEndtime: time.Unix(1575972000, 0), //现场接受资料结束时间   //2019/12/10 18:00:00
		ReceptionSite:    "长安-航天基地",                //现场接受资料地点
		LotteryData:      time.Unix(1576116000, 0), //摇号日期   2019/12/12 10:00:00
	},
	"天诚蔷薇公馆": {
		BuildID:          "201905090",              //楼盘号
		BuildName:        "天诚蔷薇公馆",                 //楼盘名
		ChooseStime:      time.Unix(1557280800, 0), //选房开始时间
		ChooseEndtime:    time.Unix(1557712800, 0), //选房结束时间    2019/5/13
		Enterprise:       "天诚中国有限公司",               //开发企业
		Introduce:        "在售小高层和洋房",               //介绍
		Hotline:          "4000386013",             //咨询电话
		PresellNumber:    "41407",                  //预售证号
		IntentStime:      time.Unix(1568340000, 0), //意向登记开始时间  2019/9/13
		IntentEndtime:    time.Unix(1568512800, 0), //意向登记结束时间
		ReceptionStime:   time.Unix(1576375200, 0), //现场接受资料开始时间  2019/12/15
		ReceptionEndtime: time.Unix(1576461600, 0), //现场接受资料结束时间
		ReceptionSite:    "辛王路与启源二路十字向东300米",       //现场接受资料地点
		LotteryData:      time.Unix(1576634400, 0), //摇号日期   2019/12/18
	},
	"远洋御山水": {
		BuildID:          "201913500",              //楼盘号
		BuildName:        "远洋御山水",                  //楼盘名
		ChooseStime:      time.Unix(1575158450, 0), //选房开始时间   2019/12/01
		ChooseEndtime:    time.Unix(1575504000, 0), //选房结束时间
		Enterprise:       "天诚中国有限公司",               //开发企业
		Introduce:        "远洋御山水积极开盘，赶快来看房吧",       //介绍
		Hotline:          "4000260153",             //咨询电话
		PresellNumber:    "41237",                  //预售证号
		IntentStime:      time.Unix(1572912000, 0), //意向登记开始时间  2019/11/05
		IntentEndtime:    time.Unix(1573171200, 0), //意向登记结束时间
		ReceptionStime:   time.Unix(1576627200, 0), //现场接受资料开始时间  2019/12/18
		ReceptionEndtime: time.Unix(1576807200, 0), //现场接受资料结束时间
		ReceptionSite:    "北辰大道",                   //现场接受资料地点
		LotteryData:      time.Unix(1577239200, 0), //摇号日期   2019/12/25
	},
}

//购房人员 type=1
var UserStore = map[string]User{
	"wangchengy": {
		ID:           0,
		Name:         "wangchengy",
		HeadImg:      "",
		Sex:          "",
		Phone:        "",
		Password:     "222222",
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
