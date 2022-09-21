/* package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 单向管道

// 只写channel
func writeChan(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("writeChan", i)
	}
	close(ch)
	wg.Done()
}

// 只读channel
func readChan(ch <-chan int) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("readChan", v)
	}
	wg.Done()
}

func main() {
	ch := make(chan int)
	wg.Add(1)
	go writeChan(ch)
	wg.Add(1)
	go readChan(ch)
	wg.Wait()
	fmt.Println("main over")
} */
