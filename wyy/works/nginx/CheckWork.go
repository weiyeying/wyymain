package nginx

import (
	"fmt"
	"time"
	"wyy/config"
	"wyy/libs"
)

func Run() {
	c := config.NginxCheck{}
	c.Nginxinit() //初始化
	check(c) //执行检测
}

func check(data config.NginxCheck) {
	//是否开启检测
	if data.IsOpen != true {
		return
	}

   checkservice(data)
}

//检测服务
func checkservice(data config.NginxCheck) {
	d := time.Duration(data.CheckTime) * time.Second
	fmt.Println(" ")
	fmt.Println("Nginx检测服务已启动......")
	for {
		datas := libs.Get(data.Url)
		if !datas {
			if data.IsSendMsg{
				libs.SendWx("nginx服务器挂掉了重新启动")
			}
			if data.IsFailedReload{
				reloadshll(data)   //杀掉服务
				startservice(data) //启动服务
			}
		}
		time.Sleep(d)
	}

}

//停止服务
func reloadshll(data config.NginxCheck)  {
	n, _ := libs.Kill(data.ServerName)
	stres:="停止服务nginx..."+ n
	libs.LogInfo(stres)

}

//启动服务
func startservice(data config.NginxCheck) {
	n,_:=libs.RunCommand(data.RestartShell)
	stre:="启动服务nginx..."+ n
	libs.LogInfo(stre)
}
