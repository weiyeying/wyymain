package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"wyy/config"
	"wyy/libs"
)

//运行
func Run() {
	//运行监测
	c := config.MysqlCheck{}
	c.Mysqlinit()
	check(c)
}

//检测服务
func check(c config.MysqlCheck) {
	if !c.IsOpen {
		return
	}
	checkmysql(c)
}

//执行检测
func checkmysql(c config.MysqlCheck) {
	d:=time.Duration(c.CheckTime)*time.Second
	fmt.Println(" ")
	fmt.Println("Mysql检测服务已启动......")
	for  {
		db, _ := sql.Open("mysql", c.Url)
		err := db.Ping()
		if err != nil {
			if c.IsFailedReload{
					killmysql(c)
					startmysql(c)
			}
			if c.IsSendMsg{
				libs.SendWx("mysql连接服务异常已重新启动")
			}

		}
		time.Sleep(d)
	}

}

//杀掉服务
func killmysql(c config.MysqlCheck) {
	n, _ := libs.Kill(c.ServerName)
	str := "重启停止服务mysql" + n
	libs.LogInfo(str)

}

//启动mysql
func startmysql(c config.MysqlCheck) {
	n,_:=libs.RunCommand(c.RestartShell)
	str:="启动mysql服务"+n
	libs.LogInfo(str)

}
