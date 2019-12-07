package server

import (
	"net/http"
	"strconv"
	"time"

	"../model"
)

var StaffHandles = RouterHandles{
	{
		Patten: "/staff/register", //员工注册
		Func:   StaffRegister,
	},
	{
		Patten: "/staff/modify", //修改员工信息
		Func:   StaffMessageModify,
	},
	{
		Patten: "/staff/delete", //删除员工信息
		Func:   DelStaff,
	},
	{
		Patten: "/staff/list", //获取员工列表
		Func:   GetStaffList,
	},
}

//员工注册
func StaffRegister(w http.ResponseWriter, r *http.Request) {
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
		ID       int   = 0
		Name           = "-"
		HeadImg        = "-" //头像URL
		Sex            = "-"
		Phone          = "-" //员工工号
		Password       = "-"
		Right          = 0 //0购房用户 1浏览 2审核 4修改 7管理员
		Status   int   = 0 //员工状态
		JoinTime int64 = 0 //员工加入时间
	)
	// 必须要的信息
	if len(r.PostForm["ID"]) > 0 &&
		len(r.PostForm["Name"]) > 0 &&
		len(r.PostForm["Sex"]) > 0 &&
		len(r.PostForm["Phone"]) > 0 &&
		len(r.PostForm["Password"]) > 0 {
		ID, _ = strconv.Atoi(r.PostForm["ID"][0])
		Name = r.PostForm["Name"][0]
		Sex = r.PostForm["Sex"][0]
		Phone = r.PostForm["Phone"][0]
		Password = r.PostForm["Password"][0]
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
	if len(r.PostForm["JoinTime"]) > 0 {
		JoinTime, _ = strconv.ParseInt(r.PostForm["RegisterTime"][0], 10, 64)
	}
	staff := model.Staff{
		ID:       ID,
		Name:     Name,
		HeadImg:  HeadImg,
		Sex:      Sex,
		Phone:    Phone,
		Password: Password,
		Right:    Right,
		Status:   Status,
		JoinTime: time.Unix(JoinTime, 0),
	}
	ok, info := staff.Register()
	if ok {
		w.Write(ReturnJsonData(0, info, "ok"))
	} else {
		w.Write(ReturnJsonData(-1, nil, "err"))
	}
}

//删除员工信息
func DelStaff(w http.ResponseWriter, r *http.Request) {
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
	if err != nil || len(r.PostForm["Name"]) <= 0 {
		w.Write(ReturnJsonData(-1, nil, "表单内容错误"))
		return
	}

	staff := new(model.Staff)
	staff.Name = r.PostForm["Name"][0]
	ok := staff.DelStaff()
	if ok {
		w.Write(ReturnJsonData(0, ok, "ok"))
	} else {
		w.Write(ReturnJsonData(-1, nil, "err"))
	}
}

//修改员工信息
func StaffMessageModify(w http.ResponseWriter, r *http.Request) {
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
		ID       int   = 0
		Name           = "-"
		HeadImg        = "-" //头像URL
		Sex            = "-"
		Phone          = "-" //员工工号
		Password       = "-"
		Right          = 0 //0购房用户 1浏览 2审核 4修改 7管理员
		Status   int   = 0 //员工状态
		JoinTime int64 = 0 //员工加入时间
	)
	// 必须要的信息
	if len(r.PostForm["ID"]) > 0 &&
		len(r.PostForm["Name"]) > 0 &&
		len(r.PostForm["Sex"]) > 0 &&
		len(r.PostForm["Phone"]) > 0 &&
		len(r.PostForm["Password"]) > 0 {
		ID, _ = strconv.Atoi(r.PostForm["ID"][0])
		Name = r.PostForm["Name"][0]
		Sex = r.PostForm["Sex"][0]
		Phone = r.PostForm["Phone"][0]
		Password = r.PostForm["Password"][0]
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
	if len(r.PostForm["JoinTime"]) > 0 {
		JoinTime, _ = strconv.ParseInt(r.PostForm["RegisterTime"][0], 10, 64)
	}
	staff := model.Staff{
		ID:       ID,
		Name:     Name,
		HeadImg:  HeadImg,
		Sex:      Sex,
		Phone:    Phone,
		Password: Password,
		Right:    Right,
		Status:   Status,
		JoinTime: time.Unix(JoinTime, 0),
	}

	ok, info := staff.StaffModifyMessage()
	if ok {
		w.Write(ReturnJsonData(0, info, "ok"))
	} else {
		w.Write(ReturnJsonData(-1, nil, "err"))
	}
}

//获取员工列表
func GetStaffList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*") // 允许跨域
	w.Header().Add("Content-type", "application/json") // 设置返回格式

	getParam := r.URL.Query() //获取URL,后面的查询参数

	if len(getParam["name"]) <= 0 || len(getParam["token"]) <= 0 {
		w.Write(ReturnJsonData(-1, nil, "参数不齐全"))
		return
	}

	staff := model.Staff{}
	w.Write(ReturnJsonData(0, staff.GetList(), "ok"))
}
