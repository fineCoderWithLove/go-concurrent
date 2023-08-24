package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total int64
var wg sync.WaitGroup

//var lock sync.Mutex

func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)
}
func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {

		atomic.AddInt64(&total, 1)  //原子性的操作

	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		atomic.AddInt64(&total,-1)
	}
}
