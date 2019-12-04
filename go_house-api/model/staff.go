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

func (u *Staff) Login() (bool, *Staff) {
	//TODO::连接数据,进行验证
	info, ok := StaffStore[u.Name]

	if !ok {
		return false, nil
	}
	if info.Name == u.Name && info.Password == u.Password {
		return true, &info
	}

	return false, nil
}

func (u *Staff) ChangeAuditStatus() {

}

//增加用户
func (u *Staff) AddUser() bool {

	return true
}

//删除用户
func (u *Staff) DelUser() bool {
	return true
}

//修改密码
func (u *Staff) ChangePsw() bool {
	return true
}

//获取权利
func (u *Staff) GetRight() (bool, int) {
	return true, u.Right
}
