package server

import "net/http"

// 把这个添加到router.go的注册方法里面
var TestHandles = RouterHandles{
	{
		Patten: "/test",
		Func:   Test,
	}, {
		Patten: "/test2",
		Func:   Test2,
	},
}

//这个方法的格式一定要对 http.ResponseWriter   *http.Request  照着复制
func Test(w http.ResponseWriter, r *http.Request) {
	w.Write(ReturnJsonData(0, "hello world!", "ok"))
}

func Test2(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"用户名": "hello",
		"密码":  "world",
	}
	w.Write(ReturnJsonData(0, data, "ok"))
}
