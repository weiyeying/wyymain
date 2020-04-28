package php

import (
	"fmt"
	"time"
	"wyy/config"
	"wyy/libs"
)

func Run()  {
	c:=config.PhpCheck{}
	c.Phpinit()
	check(c)
}

func check(c config.PhpCheck)  {
	if !c.IsOpen{
		return
	}
	phpcheck(c)

}

func phpcheck(c config.PhpCheck)  {
	d:=time.Duration(c.CheckTime)*time.Second
	fmt.Println(" ")
	fmt.Println("PHP检测服务已启动......")
	for{
		code,status:=libs.GetData(c.Url)
		if code==0{
			time.Sleep(d)
			return
		}
		if code!=200{
			libs.LogInfo("php检测故障:"+status)
			if c.IsFailedReload{
				libs.Kill(c.ServerName)
				libs.RunCommand(c.RestartShell)
				libs.LogInfo("重新启动php")
			}
			if c.IsSendMsg {
				libs.SendWx("php服务异常已重新启动"+status)
			}


		}

		time.Sleep(d)
	}

}