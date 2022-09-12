/* package main

import "fmt"

func test1() {
	var a1 [2]int
	var a2 [3]string
	fmt.Printf("a1: %T\n", a1)
	fmt.Printf("a2: %T\n", a2)
}

func test2() {
	var a1 = [4]int{1, 2}
	fmt.Printf("a1: %v\n", a1)
}

func test3() {
	var a1 = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("len(a1): %v\n", len(a1))
	var a2 = [...]int{1: 1, 3: 3, 5: 5}
	fmt.Printf("a2: %v\n", a2)
	a2[0] = 10
	fmt.Printf("a2: %v\n", a2)
	// 数组长度越界
	// a2[6] = 10
	// 数组长度
	fmt.Printf("len(a2): %v\n", len(a2))
	// 数组遍历
	for i := 0; i < len(a2); i++ {
		fmt.Printf("a2[%v]: %v\n", i, a2[i])
	}
	for _, v := range a2 {
		fmt.Printf("v: %v\n", v)
	}
}

func main() {
	/* test1()
	test2() */
	test3()
}
 */