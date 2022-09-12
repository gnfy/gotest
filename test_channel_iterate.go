/* package main

import "fmt"

var c = make(chan int, 10)

func main() {

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	for {
		v, ok := <-c
		if ok == false {
			break
		}
		fmt.Println("for v:", v)
	}
	for v := range c {
		fmt.Println("v:", v)
	}
	r := <-c
	fmt.Println("r:", r)
	r = <-c
	fmt.Println("r:", r)

}
*/