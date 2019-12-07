package server

import (
	"net/http"
	"strconv"
	"time"

	"../model"
)

var UserHandles = RouterHandles{
	//{
	//	Patten: "/user/login", //用户登陆
	//	Func:   ChangeUserAuditStatus,
	//},
	{
		Patten: "/user/register", //用户注册
		Func:   UserRegister,
	},
	{
		Patten: "/user/modify", //修改用户信息
		Func:   UserMessageModify,
	},
	{
		Patten: "/user/delete", //删除用户信息
		Func:   UserDelete,
	},
	{
		Patten: "/user/list", //获取用户列表
		Func:   GetUserList,
	},
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
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
		ID           int   = 0
		Name               = "-"
		HeadImg            = "-" //头像URL
		Sex                = "-"
		Phone              = "-"
		Password           = "-"
		Right              = 0   //0购房用户 1浏览 2审核 4修改 7管理员
		RegisterTime int64 = 0   //登记时间
		Status       int   = 0   //是否购房成功
		HouseType          = "-" //意向房源类型
		PurposeHouse       = "-" //意向房号
		IdentityType       = "-" //证件类型
		IdentityNum        = "-" //证件号码
		CheckStatus  int   = 0   //审核状态
	)

	// 必须要的信息
	if len(r.PostForm["ID"]) > 0 &&
		len(r.PostForm["Name"]) > 0 &&
		len(r.PostForm["Sex"]) > 0 &&
		len(r.PostForm["Phone"]) > 0 &&
		len(r.PostForm["Password"]) > 0 &&
		len(r.PostForm["IdentityType"]) > 0 &&
		len(r.PostForm["IdentityNum"]) > 0 {
		ID, _ = strconv.Atoi(r.PostForm["ID"][0])
		Name = r.PostForm["Name"][0]
		Sex = r.PostForm["Sex"][0]
		Phone = r.PostForm["Phone"][0]
		Password = r.PostForm["Password"][0]
		IdentityType = r.PostForm["IdentityType"][0]
		IdentityNum = r.PostForm["IdentityNum"][0]
	} else {
		w.Write(ReturnJsonData(-1, nil, "参数不完整"))
		return
	}
	//可以选择的信息
	if len(r.PostForm["HeadImg"]) > 0 {
		HeadImg = r.PostForm["HeadImg"][0]
	}
	if len(r.PostForm["Status"]) > 0 {
		Status, _ = strconv.Atoi(r.PostForm["Status"][0])
	}
	if len(r.PostForm["HouseType"]) > 0 {
		HouseType = r.PostForm["HouseType"][0]
	}
	if len(r.PostForm["PurposeHouse"]) > 0 {
		PurposeHouse = r.PostForm["PurposeHouse"][0]
	}
	if len(r.PostForm["CheckStatus"]) > 0 {
		CheckStatus, _ = strconv.Atoi(r.PostForm["CheckStatus"][0])
	}
	if len(r.PostForm["RegisterTime"]) > 0 {
		RegisterTime, _ = strconv.ParseInt(r.PostForm["RegisterTime"][0], 10, 64)
	}
	user := model.User{
		ID:           ID,
		Name:         Name,
		HeadImg:      HeadImg,
		Sex:          Sex,
		Phone:        Phone,
		Password:     Password,
		Right:        Right,
		RegisterTime: time.Unix(RegisterTime, 0),
		Status:       Status,
		HouseType:    HouseType,
		PurposeHouse: PurposeHouse,
		IdentityType: IdentityType,
		IdentityNum:  IdentityNum,
		CheckStatus:  CheckStatus,
	}
	ok, info := user.Register()
	if ok {
		w.Write(ReturnJsonData(0, info, "ok"))
	} else {
		w.Write(ReturnJsonData(-1, nil, "err"))
	}
}

func ChangeUserAuditStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*") // 允许跨域
	w.Header().Add("Content-type", "application/json") // 设置返回格式

}

//修改用户信息
func UserMessageModify(w http.ResponseWriter, r *http.Request) {
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
	user := new(model.User)

	user.ID, _ = strconv.Atoi(r.PostForm["ID"][0])
	user.Name = r.PostForm["Name"][0]
	user.HeadImg = r.PostForm["HeadImg"][0]
	user.Sex = r.PostForm["Sex"][0]
	user.Phone = r.PostForm["Phone"][0]
	user.Password = r.PostForm["Password"][0]
	user.Right, _ = strconv.Atoi(r.PostForm["Right"][0])
	var RTime int64
	RTime, _ = strconv.ParseInt(r.PostForm["RegisterTime"][0], 10, 64)
	user.RegisterTime = time.Unix(RTime, 0)
	user.Status, _ = strconv.Atoi(r.PostForm["Status"][0])
	user.HouseType = r.PostForm["HouseType"][0]
	user.PurposeHouse = r.PostForm["PurposeHouse"][0]
	user.IdentityType = r.PostForm["IdentityType"][0]
	user.IdentityNum = r.PostForm["IdentityNum"][0]
	user.CheckStatus, _ = strconv.Atoi(r.PostForm["CheckStatus"][0])

	ok, info := user.ModifyMessage()
	if ok {
		w.Write(ReturnJsonData(0, info, "ok"))
	} else {
		w.Write(ReturnJsonData(-1, nil, "err"))
	}
}

//删除用户信息
func UserDelete(w http.ResponseWriter, r *http.Request) {
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
	user := new(model.User)
	user.Name = r.PostForm["Name"][0]
	ok := user.CancelPurchase()
	if ok {
		w.Write(ReturnJsonData(0, ok, "ok"))
	} else {
		w.Write(ReturnJsonData(-1, nil, "err"))
	}
}

//获取用户列表
func GetUserList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*") // 允许跨域
	w.Header().Add("Content-type", "application/json") // 设置返回格式

	getParam := r.URL.Query() //获取URL,后面的查询参数

	if len(getParam["name"]) <= 0 || len(getParam["token"]) <= 0 {
		w.Write(ReturnJsonData(-1, nil, "参数不齐全"))
		return
	}

	user := model.User{}
	w.Write(ReturnJsonData(0, user.GetList(), "ok"))
}
