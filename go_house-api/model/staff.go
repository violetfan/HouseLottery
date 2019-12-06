package model

import (
	"fmt"
	"time"
)

//员工数据模型
type Staff struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	HeadImg     string    `json:"head_img"`
	Sex         string    `json:"sex"`
	Phone       string    `json:"phone"`
	Password    string    `json:"-"`
	Right       int       `json:"right"` // 0购房用户 1浏览 2审核 4修改 7管理员
	StaffStatus int       `json:"staff_status"`
	JoinTime    time.Time `json:"join_time"`
}

//员工登录
func (u *Staff) Login() (bool, *Staff){
	//TODO::连接数据，进行验证
	infos, ok := StaffStore[u.Name]
	fmt.Println("start Staff Login",infos)
	if !ok{
		return false,nil
	}
	if infos.Name == u.Name && infos.Password == u.Password{
		return true,&infos
	}
	return false,nil
}


func (u *Staff) ChangeAuditStatus() {

}

//注册员工
func (u *Staff) Register() (bool,*Staff) {
	if _, ok := StaffStore[u.Name];ok{
		//已注册过
		return false,nil
	}
	StaffStore[u.Name] = *u
	return true,u
}

//删除员工
func (u *Staff) DelStaff() bool {
	if _,ok := StaffStore[u.Name];ok{
		//已注册过
		delete(StaffStore,u.Name)
		return true
	}
	return false
}

//修改密码
func (u *Staff) ChangePsw() bool {
	return true
}

//获取权利
func (u *Staff) GetRight() (bool, int) {
	return true, u.Right
}
