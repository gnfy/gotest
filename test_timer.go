/* package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 5)
	fmt.Printf("now: %v\n", time.Now())
	// t1 := <-timer.C
	// fmt.Printf("t1: %v\n", t1)
	// <-timer.C
	// fmt.Printf("now: %v\n", time.Now())
	// time.Sleep(time.Second * 2)
	// fmt.Println("2s later")
	// <-time.After(time.Second)
	// fmt.Println("2s later2")
	go func() {
		<-timer.C
		fmt.Println("timer expired")
	}()
	timer.Reset(time.Second * 1)
	// stop := timer.Stop()
	// if stop {
	// 	fmt.Println("timer stopped")
	// }
	time.Sleep(time.Second * 3)
} */
