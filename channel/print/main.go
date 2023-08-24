package main

import (
	"fmt"
	"time"
)

var number, letter = make(chan bool), make(chan bool)

func printNum() {
	i := 1
	for {
		//等待另外一个goroutine进行通知
		<-number //从number进行取值的操作，如果没有值就阻塞
		fmt.Printf("%d%d", i, i+1)
		i += 2
		letter <- true
	}
}
func printLetter() {
	str := "abcdefghijklmnopqrstuvwxyz"
	i := 0
	for {
		<-letter
		if i>= len(str){
			return
		}
			fmt.Print(str[i : i+2])
		i += 2
		number <- true // 存入true到channel中
	}
}

func main() {
	go printNum()
	go printLetter()
	number <- true
	time.Sleep(10*time.Second)
}
