package config

//企业微信机器人消息推送地址申请文档https://work.weixin.qq.com/api/doc/90000/90136/91770
var Wxurl="https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=d4f81470"
var WxPhone="13311526560" //@提醒人 可为空

//nginx配置监控参数
type NginxCheck struct {
	Configs
}
func (c *NginxCheck) Nginxinit() {
	c.ServerName="nginx"  //服务名称
	c.IsOpen = true       //是否开启服务
	c.IsSendMsg=true
	c.Url = "http://dai.fenxiangtu.com/check.html" //检测地址 地址需要返回1
	c.IsFailedReload = true ////服务异常是否重新启动
	c.RestartShell = "/usr/local/nginx/sbin/nginx"  //重启服务命令
	c.CheckTime=5  //检测间隔时间（秒）
}

//mysql检测
type MysqlCheck struct {
	Configs
}

func (m *MysqlCheck) Mysqlinit()  {
	m.ServerName="mysql"  //服务名称
	m.IsOpen = true       //是否开启服务
	m.Url = "root:root@(127.0.0.1:3306)/godata" //mysql账号密码数地址端口数据库
	m.IsFailedReload = true ////服务异常是否重新启动
	m.IsSendMsg=true
	m.RestartShell = "service mysqld restart"  //重启mysql服务命令
	m.CheckTime=5  //检测间隔时间（秒）
}

