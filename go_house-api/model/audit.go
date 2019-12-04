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
