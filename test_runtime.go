/* package main

import (
	"fmt"
	"runtime"
	"time"
)

func show(msg string) {
	for i := 0; i < 2; i++ {
		fmt.Printf("i:%v, msg:%v\n", i, msg)
	}
}

func show2(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("i:%v, msg:%v\n", i, msg)
		if i >= 5 {
			runtime.Goexit()
		}
	}
}

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("a:", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("b:", i)
	}
}

func main() {
	// go show("goroutine")
	// go show2("goroutine")
	// for i := 0; i < 10; i++ {
	// 	runtime.Gosched()
	// }
	fmt.Println("CPU:", runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second * 1)
	fmt.Println("end...")
} */
