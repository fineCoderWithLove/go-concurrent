package main

import (
	"fmt"
	"time"
)

func main() {
	// 名字  类型  存储类型
	var msg chan string //默认值未nil

	msg = make(chan string, 0) //channel的初始化的值如果是0的话，放值进去会阻塞，如果设为0就为无缓冲channel
	go func(msg chan string) {
		name := <-msg //将channel中的值取出来给name
		fmt.Println(name)
	}(msg)
	msg <- "大大怪" //将右边的值放在channel中

	time.Sleep(time.Second * 10)

}
