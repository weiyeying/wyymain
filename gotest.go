package main

import (
	"fmt"
	"time"
)

func main() {
	i:=0
	for  {
		fmt.Println(i)
		time.Sleep(5*time.Second)
		i++
	}
}
