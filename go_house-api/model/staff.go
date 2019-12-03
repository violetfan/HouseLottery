package model

//员工数据模型
type Staff struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Right    int    `json:"right"` // 0购房用户 1浏览 2审核 4修改 7管理员
	// 其他信息
}

func (u *Staff) CheckPsw() (check bool, right int) {
	//TODO::连接数据,进行验证
	return true, 7
}

func (u *Staff) ChangeAuditStatus() {

}
