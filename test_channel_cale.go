/* package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f1(n int) {
	i := (n-1)*30000 + 1
	max := n * 30000
	for ; i < max; i++ {
		var flag = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			// fmt.Println(i)
		}
	}
	wg.Done()
}

func main() {
	start := time.Now().Unix()
	// max := 120000
	// for i := 2; i < max; i++ {
	// 	var flag = true
	// 	for j := 2; j < i; j++ {
	// 		if i%j == 0 {
	// 			flag = false
	// 			break
	// 		}
	// 	}
	// 	if flag {
	// 		// fmt.Println(i)
	// 	}
	// }

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i + 1)
	}
	wg.Wait()
	end := time.Now().Unix()
	fmt.Println(end - start)
	fmt.Println("over")
}
*/