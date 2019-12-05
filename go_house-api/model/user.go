package model

import (
	"fmt"
	"time"
)

// 购房用户数据模型
type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	HeadImg      string    `json:"head_img"`
	Sex          string    `json:"sex"`
	Phone        string    `json:"phone"`
	Password     string    `json:"-"`
	Right        int       `json:"right"`         // 0购房用户 1浏览 2审核 4修改 7管理员
	RegisterTime time.Time `json:"register_time"` //登记时间
	Status       int       `json:"status"`        //是否购房成功
	HouseType    string    `json:"house_type"`    //意向房源类型
	PurposeHouse string    `json:"purpose_house"` //意向房号
	IdentityType string    `json:"identity_type"` //证件类型
	IdentityNum  string    `json:"identity_num"`  //证件号码
	CheckStatus  int       `json:"check_status"`  //审核状态
}

//用户登陆
func (u *User) Login() (bool, *User) {
	//TODO::连接数据,进行验证
	infos, ok := UserStore[u.Name]
	fmt.Println("start User Login", infos)
	if !ok {
		return false, nil
	}
	if infos.Name == u.Name && infos.Password == u.Password {
		return true, &infos
	}
	return false, nil
}

//购房信息登记
func (u *User) Register() (check bool, info User) {
	return true, *u
}

//购房信息审核
func (u *User) Check() (check bool, info User) {
	return true, *u
}

//购房信息撤销
func (u *User) CancelPurchase() bool {
	return true
}

//购房登记信息修改
func (u *User) ModifyMessage() bool {
	return true
}

//TODO::增加
func (u *User) HouseRegister() {

}
