package libs

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

type System struct {
	Os string
	Numcpu int8
}



//根据进程名判断进程是否运行
func CheckProRunning(serverName string) (bool, error) {
	a := `ps ux | awk '/` + serverName + `/ && !/awk/ {print $2}'`
	pid, err := RunCommand(a)
	if err != nil {
		return false, err
	}
	return pid != "", nil
}
//根据进程名称获取进程ID
func GetPid(serverName string) (string, error) {
	a := `ps aux | awk '/` + serverName + `/ && !/awk/ {print $3}'`
	pid, err := RunCommand(a)
	return pid , err
}

//杀掉进程
func  Kill(serverName string)(string,error)  {
	a:="pkill  "+serverName
	data,err:=RunCommand(a)
	return data,err
}




func RunCommand(cmd string) (string, error) {
	if runtime.GOOS == "windows" {
		return runInWindows(cmd)
	} else {
		return runInLinux(cmd)
	}
}
func runInWindows(cmd string) (string, error) {
	result, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(result)), err
}

func runInLinux(cmd string) (string, error) {
	fmt.Println("Running Linux cmd:" + cmd)
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return strings.TrimSpace(string(result)), err
}