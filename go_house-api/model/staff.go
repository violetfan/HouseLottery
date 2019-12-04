package model

import "time"

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

func (u *Staff) Login() (check bool, info Staff) {
	//TODO::连接数据,进行验证
	if u.Name == "admin" && u.Password == "1234567" {
		return true, *u
	}
	return true, *u
}

func (u *Staff) ChangeAuditStatus() {

}

func (u *Staff) AddUser() {

}
