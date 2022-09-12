/* package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var i int = 100

var lock sync.Mutex

func add() {
	lock.Lock()
	i++
	lock.Unlock()
}

func sub() {
	lock.Lock()
	i--
	lock.Unlock()
}

var a int32 = 100

func add2() {
	atomic.AddInt32(&a, 1)
}

func sub2() {
	atomic.AddInt32(&a, -1)
}

var b int32 = 100

func load_store() {
	atomic.LoadInt32(&b)
	fmt.Println("b:", b)
	atomic.StoreInt32(&b, 200)
	fmt.Println("b:", b)
}

func compareAndSwap() {
	c := atomic.CompareAndSwapInt32(&b, 200, 201)
	if c {
		fmt.Println("b true:", b)
	} else {
		fmt.Println("b false:", b)
	}
}

func main() {
	for i := 0; i < 1000; i++ {
		go add()
		go sub()
	}
	for i := 0; i < 1000; i++ {
		go add2()
		go sub2()
	}
	time.Sleep(time.Second * 2)
	fmt.Println("i:", i)
	fmt.Println("a:", a)
	load_store()
	compareAndSwap()
}
*/