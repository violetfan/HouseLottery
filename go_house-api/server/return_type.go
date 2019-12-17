package server

import "encoding/json"

type ReturnType struct {
	ErrCode int         `json:"err_code"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
}

func Data2json(data interface{}) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		println(err.Error())
		return []byte("")
	} else {
		return bytes
	}

}

func ReturnJsonData(errCode int, data interface{}, msg string) []byte {
	returndata := ReturnType{ErrCode: errCode, Data: data, Msg: msg}
	return Data2json(returndata)
}
