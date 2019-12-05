package server

import "net/http"

type RouterHandle struct {
	Patten string
	Func   http.HandlerFunc
}

type RouterHandles []*RouterHandle

func RegisterRouter(mux *http.ServeMux) {
	handles := RouterHandles{}

	// 在这里添加路由
	handles = append(handles, LoginHandles...) //登入路由
	handles = append(handles, UserHandles...)  //用户路由
	handles = append(handles, HouseHandles...) //房源路由

	//test
	handles = append(handles, TestHandles...)

	for _, router := range handles {
		mux.HandleFunc(router.Patten, router.Func)
	}
}
