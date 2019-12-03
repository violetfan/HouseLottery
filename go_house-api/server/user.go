package server

import "net/http"

var UserHandles = RouterHandles{
	{
		Patten: "/user/login",
		Func:   ChangeUserAuditStatus,
	},
}

func CheckUserPsw() {

}

func ChangeUserAuditStatus(w http.ResponseWriter, r *http.Request) {

}
