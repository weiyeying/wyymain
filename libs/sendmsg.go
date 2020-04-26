package libs

import (
	"wyy/config"
)

//发送微信消息  url地址  msg 消息内容
func SendWx(msg string, phone ...string) {
	url := config.Wxurl
	wphone:=config.WxPhone
	maps := make(map[string]interface{})
	maps["msgtype"] = "text"
	maps2 := make(map[string]interface{})
	maps2["content"] = msg
	phonelist := []string{}
	for _, v := range phone {
		phonelist = append(phonelist, v)
	}
	phonelist=append(phonelist,wphone)
	maps2["mentioned_mobile_list"] = phonelist
	maps["text"] = maps2
	Post(url, maps, "application/json")

}
