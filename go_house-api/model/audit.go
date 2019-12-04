package model

import "time"

//审核记录表
type AuditNotes struct {
	ID          int       `json:"id"`
	AuditTime   time.Time `json:"audit_time"`
	StaffID     string    `json:"staff_id"`
	AuditObject string    `json:"audit_object"`
	AuditStatus string    `json:"audit_status"`
	AuditID     int       `json:"audit_id"`
}

//摇号结果
type LotteryResult struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	PaperNum string `json:"paper_num"`
}

//选房结果
type HouseResult struct {
	HouseNum   string    `json:"house_num"`
	HouseState bool      `json:"house_state"`
	TimeStamp  time.Time `json:"time_stamp"`
}

//审核信息类
//审核 增加 ，obj：用户u_id/员工s_id
func (a *AuditNotes) AddRecord(obj string) (bool, *AuditNotes) {
	return true, a
}

//审核 撤销，obj：用户u_id/员工s_id
func (a *AuditNotes) DelRecord(obj string) (bool, *AuditNotes) {
	return true, a
}

//修改审核状态
func (a *AuditNotes) ChangeStatus(obj string) (bool, *AuditNotes) {
	return true, a
}

//导出数据
func (a *AuditNotes) Output(obj string) (bool, *AuditNotes) {
	return true, a
}

//摇号信息类
//增加摇号
func (l *LotteryResult) AddSelect() bool {
	return true
}

//删除摇号
func (l *LotteryResult) DelSelect() bool {
	return true
}

//选房结果类
//签字确认
func (h *HouseResult) Confirm() bool {
	return true
}

//信息隐藏
func (h *HouseResult) Hide() bool {
	return true
}

//信息显示
func (h *HouseResult) Show() (bool, *HouseResult) {
	return true, h
}
