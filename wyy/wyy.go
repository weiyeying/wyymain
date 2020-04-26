package wyy

import (
	"fmt"
	"os"
	"wyy/libs"
	"wyy/wyy/works/mysql"
	"wyy/wyy/works/nginx"
)

func Run() {
	arg := os.Args
	if len(arg) < 2 {
		fmt.Println("[err]参数错误 运行时请添加参数 [start]开启 [stop ]关闭...")
		return
	}
	status := arg[1]
	if status == "start" {
		start()
	}
	if status == "stop" {
		fmt.Println("[stop]服务停止...")
		stop()
	}
	if status != "start" && status != "stop" {
		fmt.Println("[err]参数错误 运行时请添加参数 [start]开启 [stop ]关闭...")
	}
	select {}
}

//注册检测模块
func Registered()  {
	go nginx.Run() //注册nginx检测
	go mysql.Run() //注册mysql检测

}

//启动服务
func start() {
	Registered() //注册

}

func stop() {
	_, err := libs.Kill("wyymain")
	if err != nil {
		fmt.Println("[err]停止失败 请手动Kill进程...")
	}
}
