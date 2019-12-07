package server

import (
	"net/http"
	"strconv"
	"time"

	"../model"
)

var HouseHandles = RouterHandles{
	{
		Patten: "/house/issue",
		Func:   IssueHouse,
	},
	{
		Patten: "/house/list",
		Func:   GetHouseList,
	},
	{
		Patten: "/house/modify",
		Func:   ModifyHouse,
	},
}

//发布房源
func IssueHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*") // 允许跨域
	w.Header().Add("Content-type", "application/json") // 设置返回格式

	if r.Method != "POST" {
		w.Write(ReturnJsonData(-1, nil, "请求方式错误"))
		return
	}

	getParam := r.URL.Query() //获取URL,后面的查询参数

	if len(getParam["name"]) <= 0 || len(getParam["token"]) <= 0 {
		w.Write(ReturnJsonData(-1, nil, "参数不齐全"))
		return
	}

	err := r.ParseForm() //解析POST参数
	if err != nil {
		w.Write(ReturnJsonData(-1, nil, "表单内容错误"))
		return
	}

	//初始化默认信息
	var (
		BuildID                = "-"
		BuildName              = "-"
		Introduce              = "-"
		Enterprise             = "-"
		Hotline                = "-"
		PresellNumber          = "-"
		ReceptionSite          = "暂无"
		ChooseStime      int64 = 0
		ChooseEndtime    int64 = 0
		IntentStime      int64 = 0
		IntentEndtime    int64 = 0
		ReceptionStime   int64 = 0
		ReceptionEndtime int64 = 0
		LotteryData      int64 = 0
	)

	// 必须要的信息
	if len(r.PostForm["BuildID"]) > 0 &&
		len(r.PostForm["BuildName"]) > 0 &&
		len(r.PostForm["Introduce"]) > 0 &&
		len(r.PostForm["Enterprise"]) > 0 &&
		len(r.PostForm["Hotline"]) > 0 &&
		len(r.PostForm["PresellNumber"]) > 0 {
		BuildID = r.PostForm["BuildID"][0]
		BuildName = r.PostForm["BuildName"][0]
		Introduce = r.PostForm["Introduce"][0]
		Enterprise = r.PostForm["Enterprise"][0]
		Hotline = r.PostForm["Hotline"][0]
		PresellNumber = r.PostForm["PresellNumber"][0]
	} else {
		w.Write(ReturnJsonData(-1, nil, "参数不完整"))
		return
	}

	//可以选择的信息
	if len(r.PostForm["ChooseStime"]) > 0 {
		ChooseStime, _ = strconv.ParseInt(r.PostForm["ChooseStime"][0], 10, 64)
	}
	if len(r.PostForm["ChooseEndtime"]) > 0 {
		ChooseEndtime, _ = strconv.ParseInt(r.PostForm["ChooseEndtime"][0], 10, 64)
	}
	if len(r.PostForm["IntentStime"]) > 0 {
		IntentStime, _ = strconv.ParseInt(r.PostForm["IntentStime"][0], 10, 64)
	}
	if len(r.PostForm["IntentEndtime"]) > 0 {
		IntentEndtime, _ = strconv.ParseInt(r.PostForm["IntentEndtime"][0], 10, 64)
	}
	if len(r.PostForm["ReceptionStime"]) > 0 {
		ReceptionStime, _ = strconv.ParseInt(r.PostForm["ReceptionStime"][0], 10, 64)
	}
	if len(r.PostForm["ReceptionEndtime"]) > 0 {
		ReceptionEndtime, _ = strconv.ParseInt(r.PostForm["ReceptionEndtime"][0], 10, 64)
	}
	if len(r.PostForm["LotteryData"]) > 0 {
		LotteryData, _ = strconv.ParseInt(r.PostForm["LotteryData"][0], 10, 64)
	}
	if len(r.PostForm["ReceptionSite"]) > 0 {
		ReceptionSite = r.PostForm["ReceptionSite"][0]
	}

	house := model.House{
		BuildID:          BuildID,
		BuildName:        BuildName,
		ChooseStime:      time.Unix(ChooseStime, 0),
		ChooseEndtime:    time.Unix(ChooseEndtime, 0),
		Enterprise:       Enterprise,
		Introduce:        Introduce,
		Hotline:          Hotline,
		PresellNumber:    PresellNumber,
		IntentStime:      time.Unix(IntentStime, 0),
		IntentEndtime:    time.Unix(IntentEndtime, 0),
		ReceptionStime:   time.Unix(ReceptionStime, 0),
		ReceptionEndtime: time.Unix(ReceptionEndtime, 0),
		ReceptionSite:    ReceptionSite,
		LotteryData:      time.Unix(LotteryData, 0),
	}

	ok, info := house.Issue()
	if ok {
		w.Write(ReturnJsonData(0, info, "ok"))
	} else {
		w.Write(ReturnJsonData(-1, nil, "错误"))
	}
}

//获取房源列表
func GetHouseList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*") // 允许跨域
	w.Header().Add("Content-type", "application/json") // 设置返回格式

	house := new(model.House)
	w.Write(ReturnJsonData(-0, house.GetHouseList(), "ok"))
}

//修改房屋信息
func ModifyHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*") // 允许跨域
	w.Header().Add("Content-type", "application/json") // 设置返回格式

	if r.Method != "POST" {
		w.Write(ReturnJsonData(-1, nil, "请求方式错误"))
		return
	}

	getParam := r.URL.Query() //获取URL,后面的查询参数

	if len(getParam["name"]) <= 0 || len(getParam["token"]) <= 0 {
		w.Write(ReturnJsonData(-1, nil, "参数不齐全"))
		return
	}

	err := r.ParseForm() //解析POST参数
	if err != nil {
		w.Write(ReturnJsonData(-1, nil, "表单内容错误"))
		return
	}
	house := new(model.House)
	house.BuildID = r.PostForm["BuildID"][0]
	house.BuildName = r.PostForm["BuildID"][0]
	CSTime, _ := strconv.ParseInt(r.PostForm["ChooseStime"][0], 10, 64)
	house.ChooseStime = time.Unix(CSTime, 0)
	CETime, _ := strconv.ParseInt(r.PostForm["ChooseEndtime"][0], 10, 64)
	house.ChooseEndtime = time.Unix(CETime, 0)
	house.Enterprise = r.PostForm["Enterprise"][0]
	house.Introduce = r.PostForm["Introduce"][0]
	house.Hotline = r.PostForm["Hotline"][0]
	house.PresellNumber = r.PostForm["PresellNumber"][0]
	ISTime, _ := strconv.ParseInt(r.PostForm["IntentStime"][0], 10, 64)
	house.IntentStime = time.Unix(ISTime, 0)
	IETime, _ := strconv.ParseInt(r.PostForm["IntentEndtime"][0], 10, 64)
	house.IntentEndtime = time.Unix(IETime, 0)
	RSTime, _ := strconv.ParseInt(r.PostForm["ReceptionStime"][0], 10, 64)
	house.ReceptionStime = time.Unix(RSTime, 0)
	RETime, _ := strconv.ParseInt(r.PostForm["ReceptionEndtime"][0], 10, 64)
	house.ReceptionEndtime = time.Unix(RETime, 0)
	house.ReceptionSite = r.PostForm["ReceptionSite"][0]
	HLTime, _ := strconv.ParseInt(r.PostForm["LotteryData"][0], 10, 64)
	house.LotteryData = time.Unix(HLTime, 0)
	ok, info := house.ModifyHouseMessage()
	if ok {
		w.Write(ReturnJsonData(0, info, "ok"))
	} else {
		w.Write(ReturnJsonData(-1, nil, "错误"))
	}
}
