/* package main

import (
	"fmt"
	"math/rand"
	"time"
)

var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixMicro())
	value := rand.Intn(10)
	fmt.Printf("send:%v\n", value)
	values <- value
}

func main() {
	defer close(values)
	go send()
	fmt.Println("wait...")
	value := <-values
	fmt.Printf("receive:%v\n", value)
	fmt.Println("end")
}
*/