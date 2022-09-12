/* package main

import (
	"fmt"
	"sync"
)

var i int = 100
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	defer wg.Done()
	lock.Lock()
	i += 1
	fmt.Println("add:", i)
	lock.Unlock()
}

func sub() {
	defer wg.Done()
	lock.Lock()
	i -= 1
	fmt.Println("sub:", i)
	lock.Unlock()
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go add()
		wg.Add(1)
		go sub()
	}
	fmt.Println("end:", i)
}
*/