package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int)  {
	for i:=0;i<10 ;i++  {
		out <- i*i
	}
	close(out)
}

func consumer(in <-chan int)  {
	for num := range in {
		fmt.Println(num)
	}
}
func main() {
	//var ch1 chan int //双向的channel
	//var ch2 chan<- float64 //单项channel,只能写入float64的数据
	//var ch3 <-chan int //只能读取int类型的数据

	/*
	定义一个channel然后把它编程单向的,但是不能把单项的channel转成双向的channel
	 */
	//c := make(chan int, 3)
	//var send chan<- int = c
	//var read <-chan int = c
	//
	//send <- 1
	//num := <- read
	//
	//fmt.Println(num)
	/*
	内部会完成自动的类型转换
	 */
	c := make(chan int)
	go producer(c)
	go consumer(c)
	time.Sleep(10*time.Second)
}
