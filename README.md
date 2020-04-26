wyymain服务（php+mysql+nginx）管理
====



一、说明
----
WYYmain是一款轻量级服务管理系统，go语言开发，部署超级简单，资源消耗少，运行稳定。
支持定时检测linux 中服务的状态 目前支持nginx mysql等服务监测

支持自定义定时检测时间



码云地址：https://gitee.com/wyymain/wyymain
Github地址:https://github.com/weiyeying/wyymain





二、特性
----


- 1、开箱即用简便的命令行操作
- 2、定时监控mysql nginx php 服务可用性 
- 3、自动重启功能 监测服务不可用自动重启
- 4、新增消息通知功能 支持企业微信消息推送




三、支持
----
1、给项目一个star





五、安装方法
----

获取安装

- go get github.com/weiyeying/wyymain
- go get github.com/go-sql-driver/mysql

- 修改config 配置文件
- 启动服务 go run wyymain.go start &
- 停止服务 go run wyymain.go stop




配置文件
----
根据自己的情况修改
```

# 服务名称
ServerName = nginx  服务名称 每个服务都是一个struct


#通知目前支持企业微信（地址可在企业微信申请）
Wxurl = https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=d4f81470  
WxPhone="13311526560"


# 配置文件
ServerName="nginx"  //服务名称
IsOpen = true       //是否开启服务
IsSendMsg=true
Url = "http://dai.fenxiangtu.com/check.html" //检测地址 地址需要返回1
IsFailedReload = true ////服务异常是否重新启动
RestartShell = "/usr/local/nginx/sbin/nginx"  //重启服务命令
CheckTime=5  //检测间隔时间（秒）

```


六、联系我
----
wx号:qq448520051

欢迎交流，欢迎提交代码。






