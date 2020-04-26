package config

type Configs struct {
	ServerName   string  //进程名称(需要与linux 进程名相同否则kill重启进程false)
	IsOpen         bool   //是否开启
	Url            string //检测地址
	IsSendMsg      bool   //服务异常是否开启通知
	IsFailedReload bool   //服务异常是否重新启动
	RestartShell   string //服务重启命令
	CheckTime int64   //检测间隔时间
}