package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// select 读取channel里面的数据，不需要关闭channel
	for {
		select {
		case v := <-intChan:
			fmt.Printf("intchan: %v\n", v)
			time.Sleep(time.Second * 1)
		case v := <-strChan:
			fmt.Printf("strchan: %v\n", v)
			time.Sleep(time.Second * 2)
		default:
			println("default")
			return
		}
	}
}
