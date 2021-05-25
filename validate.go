package main

import (
	"RMQ_Project/common"
	"fmt"
	"net/http"
)

func Auth(rw http.ResponseWriter, rq *http.Request) error {
	return nil
}
func Check(rw http.ResponseWriter, rq *http.Request)  {
	fmt.Print("执行check")
}


func main() {
	// 1 过滤器
	filter := common.NewFilterHandle()
	// 注册拦截器
	filter.RegisterFilterUri("/check", Auth)
	// 2 启动服务
	http.HandleFunc("/check", filter.Handle(Check))
	//监听端口
	_ = http.ListenAndServe(":8081", nil)
}
