package server

var StaffHandles = RouterHandles{
	{
		Patten: "/admin/login",
		Func:   ChangeUserAuditStatus,
	},
}

func CheckStaffPsw() {

}
