package model

import (
	"fmt"
	"time"
)

//员工数据模型
type Staff struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	HeadImg  string    `json:"head_img"`
	Sex      string    `json:"sex"`
	Phone    string    `json:"phone"`
	Password string    `json:"-"`
	Right    int       `json:"right"` // 0购房用户 1浏览 2审核 4修改 7管理员
	Status   int       `json:"status"`
	JoinTime time.Time `json:"join_time"`
}

//员工登录
func (u *Staff) Login() (bool, *Staff) {
	infos, ok := StaffStore[u.Name]
	fmt.Println("start Staff Login", infos)
	if !ok {
		return false, nil
	}
	if infos.Name == u.Name && infos.Password == u.Password {
		return true, &infos
	}
	return false, nil
}

func (u *Staff) GetList() (StaffList []*Staff) {
	for _, v := range StaffStore {
		StaffList = append(StaffList, &v)
	}
	return StaffList
}

//注册员工
func (u *Staff) Register() (bool, *Staff) {
	if _, ok := StaffStore[u.Name]; ok {
		//已注册过
		return false, nil
	}
	StaffStore[u.Name] = *u
	return true, u
}

//删除员工
func (u *Staff) DelStaff() bool {
	if _, ok := StaffStore[u.Name]; ok {
		//已注册过
		delete(StaffStore, u.Name)
		return true
	}
	return false
}

//员工登记信息修改
func (u *Staff) StaffModifyMessage() (bool, *Staff) {
	if _, ok := StaffStore[u.Name]; ok { //已登记过
		StaffStore[u.Name] = *u
		return true, u
	}
	return false, nil
}

//修改密码
func (u *Staff) ChangePsw() bool {
	return true
}

//获取权利
func (u *Staff) GetRight() (bool, int) {
	return true, u.Right
}
