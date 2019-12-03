package model

// 购房用户数据模型
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Right    int    `json:"right"` // 0购房用户 1浏览 2审核 4修改 7管理员
	// 其他信息
}

func (u *User) CheckPsw() (check bool) {
	//TODO::连接数据,进行验证
	return true
}

//TODO::增加
func (u *User) HouseRegister() {

}
