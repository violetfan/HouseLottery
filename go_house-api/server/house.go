package server

import (
	"fmt"
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
	{
		Patten: "/house/del",
		Func:   DelHouse,
	},
}

//发布房源
func IssueHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json") // 设置返回格式
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Access-Token")
	w.Header().Add("Access-Control-Expose-Headers", "*")

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
		PicView                = ""
		ChooseEndtime    int64 = 0
		IntentStime      int64 = 0
		IntentEndtime    int64 = 0
		ReceptionStime   int64 = 0
		ReceptionEndtime int64 = 0
		LotteryDate      int64 = 0
	)
	// 必须要的信息
	if len(r.PostForm["build_id"]) > 0 &&
		len(r.PostForm["build_name"]) > 0 &&
		len(r.PostForm["introudce"]) > 0 &&
		len(r.PostForm["enterprise"]) > 0 &&
		len(r.PostForm["hotline"]) > 0 &&
		len(r.PostForm["hotline"]) > 0 &&
		len(r.PostForm["presell_number"]) > 0 {
		BuildID = r.PostForm["build_id"][0]
		PicView = r.PostForm["pic_view"][0]
		BuildName = r.PostForm["build_name"][0]
		Introduce = r.PostForm["introudce"][0]
		Enterprise = r.PostForm["enterprise"][0]
		Hotline = r.PostForm["hotline"][0]
		PresellNumber = r.PostForm["presell_number"][0]
	} else {
		w.Write(ReturnJsonData(-1, nil, "参数不完整"))
		return
	}

	//可以选择的信息
	if len(r.PostForm["choose_stime"]) > 0 {
		ChooseStime, _ = strconv.ParseInt(r.PostForm["choose_stime"][0], 10, 64)
	}
	if len(r.PostForm["choose_endtime"]) > 0 {
		ChooseEndtime, _ = strconv.ParseInt(r.PostForm["choose_endtime"][0], 10, 64)
	}
	if len(r.PostForm["intent_stime"]) > 0 {
		IntentStime, _ = strconv.ParseInt(r.PostForm["intent_stime"][0], 10, 64)
	}
	if len(r.PostForm["intent_endtime"]) > 0 {
		IntentEndtime, _ = strconv.ParseInt(r.PostForm["intent_endtime"][0], 10, 64)
	}
	if len(r.PostForm["reception_stime"]) > 0 {
		ReceptionStime, _ = strconv.ParseInt(r.PostForm["reception_stime"][0], 10, 64)
	}
	if len(r.PostForm["reception_endtime"]) > 0 {
		ReceptionEndtime, _ = strconv.ParseInt(r.PostForm["reception_endtime"][0], 10, 64)
	}
	if len(r.PostForm["lottery_date"]) > 0 {
		LotteryDate, _ = strconv.ParseInt(r.PostForm["lottery_date"][0], 10, 64)
	}
	if len(r.PostForm["reception_site"]) > 0 {
		ReceptionSite = r.PostForm["reception_site"][0]
	}

	house := model.House{
		BuildID:          BuildID,
		BuildName:        BuildName,
		ChooseStime:      time.Unix(ChooseStime, 0),
		ChooseEndtime:    time.Unix(ChooseEndtime, 0),
		Enterprise:       Enterprise,
		Introduce:        Introduce,
		Hotline:          Hotline,
		PicView:          PicView,
		PresellNumber:    PresellNumber,
		IntentStime:      time.Unix(IntentStime, 0),
		IntentEndtime:    time.Unix(IntentEndtime, 0),
		ReceptionStime:   time.Unix(ReceptionStime, 0),
		ReceptionEndtime: time.Unix(ReceptionEndtime, 0),
		ReceptionSite:    ReceptionSite,
		LotteryDate:      time.Unix(LotteryDate, 0),
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

//删除房源
func DelHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json") // 设置返回格式
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Access-Token")
	w.Header().Add("Access-Control-Expose-Headers", "*")

	getParam := r.URL.Query() //获取URL,后面的查询参数

	if len(getParam["name"]) <= 0 || len(getParam["token"]) <= 0 {
		w.Write(ReturnJsonData(-1, nil, "参数不齐全"))
		return
	}

	houseName := getParam["build_name[]"]
	for _, v := range houseName {
		fmt.Println(v)
		hous := model.House{BuildName: v}
		hous.Del()
	}

	w.Write(ReturnJsonData(0, nil, "ok"))
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
	HLTime, _ := strconv.ParseInt(r.PostForm["LotteryDate"][0], 10, 64)
	house.LotteryDate = time.Unix(HLTime, 0)
	ok, info := house.ModifyHouseMessage()
	if ok {
		w.Write(ReturnJsonData(0, info, "ok"))
	} else {
		w.Write(ReturnJsonData(-1, nil, "错误"))
	}
}
