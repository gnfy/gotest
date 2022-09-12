/* package main

import "fmt"

func main() {
	var ip *int
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("ip: %T\n", ip)
	var i int = 100
	ip = &i
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("ip: %v\n", *ip)
	a := [4]int{1, 2, 3, 4}
	var ptr [4]*int
	for i := 0; i < len(a); i++ {
		ptr[i] = &a[i]
	}
	fmt.Printf("ptr: %v\n", ptr)
	for i := 0; i < len(ptr); i++ {
		fmt.Printf("ptr:%v\n", *ptr[i])
	}
}
*/