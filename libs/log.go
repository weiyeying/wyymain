package libs

import (
	"io"
	"os"
	"time"
)

const (
	LOGPATH = "log/"
	FORMAT = "20060102"
	FORMATS="2006-01-02 15:04:05"
	LineFeed = "\r\n"
)

var path = LOGPATH

func LogInfo(msg string) error {
	if !IsExist(path) {
		return CreateDir(path)
	}
	var (
		err error
		f   *os.File
	)
	fileName := path + time.Now().Format(FORMAT) + ".log"
	f, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	msg="[info]"+time.Now().Format(FORMATS)+":" +msg
	_, err = io.WriteString(f, LineFeed+msg)
	defer f.Close()
	return err
}

//CreateDir  文件夹创建
func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	os.Chmod(path, os.ModePerm)
	return nil
}

//IsExist  判断文件夹/文件是否存在  存在返回 true
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
