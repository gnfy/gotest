package main

import (
	"fmt"
	"time"
)

func show(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("i:%v, msg:%v\n", i, msg)
		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {
	go show("goroutine")
	show("go")
}
