package model

import "time"

//售楼公告
type House struct {
	BuildID          string    `json:"build_id"`
	BuildName        string    `json:"build_name"`
	ChooseStime      time.Time `json:"choose_stime"`
	ChooseEndtime    time.Time `json:"choose_endtime"`
	Enterprise       string    `json:"enterprise"`
	Introduce        string    `json:"introudce"`
	Hotline          string    `json:"hotline"`
	PicView          string    `json:"pic_view"`
	PresellNumber    string    `json:"presell_number"`
	IntentStime      time.Time `json:"intent_stime"`
	IntentEndtime    time.Time `json:"intent_endtime"`
	ReceptionStime   time.Time `json:"reception_stime"`
	ReceptionEndtime time.Time `json:"reception_endtime"`
	ReceptionSite    string    `json:"reception_site"`
	LotteryDate      time.Time `json:"lottery_date"`
}

//审核
func (h *House) Audit() bool {
	return true
}

//删除
func (h *House) Del() {
	delete(HouseStore, h.BuildName)
}

//发布
func (h *House) Issue() (bool, *House) {
	//TODO::验证信息是否在正确,是否存在同名  BuildID ? BuildName ,统一一下主键
	HouseStore[h.BuildName] = *h
	return true, h
}

//获取房源
func (h *House) GetHouseList() (HouseList []House) {
	for _, v := range HouseStore {
		HouseList = append(HouseList, v)
	}
	return HouseList
}

//修改房屋信息
func (h *House) ModifyHouseMessage() (bool, *House) {
	if _, ok := HouseStore[h.BuildName]; ok { //已登记过
		HouseStore[h.BuildName] = *h
		return true, h
	}
	return false, h
}
