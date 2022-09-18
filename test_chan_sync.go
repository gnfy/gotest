/* package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func f1(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("写入数据：", i)
	}
	close(ch)
	wg.Done()
}

func f2(ch chan int) {
	// for {
	// 	v, ok := <-ch
	// 	if !ok {
	// 		fmt.Println("读取完毕")
	// 		break
	// 	}
	// 	fmt.Println("读取数据：", v)
	// }
	for v := range ch {
		fmt.Printf("读取数据: %v\n", v)
	}
	wg.Done()
}

func main() {
	ch := make(chan int, 10)
	wg.Add(1)
	go f1(ch)
	wg.Add(1)
	go f2(ch)
	wg.Wait()
	fmt.Println("over")
}
*/