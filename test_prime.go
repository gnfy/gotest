/* package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 放入数据进channel
func putData(intChan chan int) {
	for i := 1; i <= 100; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}

// 从channel取出数据,并判断是否为素数
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	// 什么时间关闭chan
	// close(primeChan)
	exitChan <- true
	wg.Done()
}

// 打印素数
func printPrime(primeChan chan int) {
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		println(res)
	}
	wg.Done()
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 10)
	wg.Add(1)
	go putData(intChan)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}
	wg.Add(1)
	go printPrime(primeChan)

	// 判断exitChan中是否有10个true
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			<-exitChan
		}
		close(primeChan)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("over")
}
*/