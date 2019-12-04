package server

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"../model"
)

var LoginHandles = RouterHandles{
	{
		Patten: "/login",
		Func:   Login,
	}, {
		Patten: "/checktoken",
		Func:   CheckToken,
	},
}

func Login(w http.ResponseWriter, r *http.Request) {
	getParam := r.URL.Query() //获取get参数

	if len(getParam["name"]) <= 0 || len(getParam["password"]) <= 0 || len(getParam["type"]) <= 0 {
		w.Write(ReturnJsonData(-1, nil, "参数不齐全"))
		return
	}

	uname := getParam["name"][0]
	upsw := getParam["password"][0]
	utype := getParam["type"][0] //0用户 1员工

	switch utype {
	case "0": //员工登入
		check, info := checkStaffPsw(uname, upsw)
		if check {
			loginfo := RecordLoginStatus(*info)
			w.Write(ReturnJsonData(-1, loginfo, "员工登入成功"))
		} else {
			w.Write(ReturnJsonData(-1, nil, "账号或者密码错误"))
		}
		break
	case "1": //用户登入
		check, info := checkUserPsw(uname, upsw)
		if check {
			loginfo := RecordLoginStatus(*info)
			w.Write(ReturnJsonData(-1, loginfo, "用户登入成功"))
		} else {
			w.Write(ReturnJsonData(-1, nil, "账号或者密码错误"))
		}
		break
	default:
		w.Write(ReturnJsonData(-12, "错误", "类型错误"))
	}
}

func CheckToken(w http.ResponseWriter, r *http.Request) {
	getParam := r.URL.Query() //获取get参数
	if len(getParam["name"]) <= 0 || len(getParam["token"]) <= 0 {
		w.Write(ReturnJsonData(-1, nil, "参数不齐全"))
		return
	}
	uname := getParam["name"][0]
	utoken := getParam["token"][0]

	info, ok := CheckLoginStatus(uname, utoken)
	if ok {
		w.Write(ReturnJsonData(0, info, "ok"))
	} else {
		w.Write(ReturnJsonData(-2, info, "token错误,或已过期"))
	}
}

func checkStaffPsw(name string, psw string) (bool, *LoginInfo) {
	//调用数据模型
	staff := model.Staff{
		Name:     name,
		Password: psw,
	}
	check, info := staff.Login()

	if check {
		logininfo := LoginInfo{
			ID:      info.ID,
			Name:    info.Name,
			HeadImg: info.HeadImg,
			Sex:     info.Sex,
			Phone:   info.Phone,
			Right:   info.Right,
		}
		return true, &logininfo
	} else {
		return false, nil
	}

}

func checkUserPsw(name string, psw string) (bool, *LoginInfo) {
	user := model.User{
		Name:     name,
		Password: psw,
	}
	check, info := user.Login()

	if check {
		logininfo := LoginInfo{
			ID:      info.ID,
			Name:    info.Name,
			HeadImg: info.HeadImg,
			Sex:     info.Sex,
			Phone:   info.Phone,
			Right:   info.Right,
		}
		return true, &logininfo
	} else {
		return false, nil
	}
}

// 其他操作校验token的函数
func CheckLoginStatus(name string, token string) (UserInfo *LoginInfo, check bool) {
	userInfo, ok := LoginInfos[token]
	if ok { //有记录
		if name != userInfo.Name {
			//校验用户名与token不符
			return nil, false
		}
		if userInfo.Expire.Unix() <= time.Now().Unix() {
			//过期
			delete(LoginInfos, token)
			return nil, false
		} else {
			//为expire续期
			userInfo.Expire = time.Now().Add(5 * time.Minute)
			LoginInfos[token] = userInfo
			return &userInfo, true
		}
	} else {
		return nil, false
	}
}

//记录登入状态
func RecordLoginStatus(info LoginInfo) LoginInfo {
	userInfo := LoginInfo{
		ID:      info.ID,
		Name:    info.Name,
		HeadImg: info.HeadImg,
		Sex:     info.Sex,
		Phone:   info.Phone,
		Right:   info.Right,
		Token:   RandToken(),
		Expire:  time.Now().Add(5 * time.Minute),
	}
	LoginInfos[userInfo.Token] = userInfo
	return userInfo
}

// 记录登入状态  (线程不安全,就不加锁了)
var LoginInfos = map[string]LoginInfo{}

type LoginInfo struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	HeadImg string    `json:"head_img"`
	Sex     string    `json:"sex"`
	Phone   string    `json:"phone"`
	Right   int       `json:"right"` // 0购房用户 1浏览 2审核 4修改 7管理员
	Token   string    `json:"token"`
	Expire  time.Time `json:"-"`
}

// 生成num*2位的字符串
func RandToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	token := fmt.Sprintf("%x", b)
	return token
}
