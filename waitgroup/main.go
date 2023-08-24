package main

import (
	"fmt"
	"sync"
)

func main()  {
	var wq sync.WaitGroup

	wq.Add(100) //监控多少个goroutine执行结束

	for i:= 0;i<100;i++ {
		go func(i int) {
			defer wq.Done()
			fmt.Println(i)
		}(i)
	}

	wq.Wait()
}

