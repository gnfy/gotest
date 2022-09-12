/* package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3, 4, 5}
	l := len(s1)
	for i := 0; i < l; i++ {
		fmt.Printf("s1[%v]: %v\n", i, s1[i])
	}
	s1 = append(s1, 6)
	for _, v := range s1 {
		fmt.Printf("v: %v\n", v)
	}
	// 删除 i = 2 的元素
	s1 = append(s1[:2], s1[3:]...)
	fmt.Printf("s1: %v\n", s1)
	// 数组复制
	s2 := make([]int, 3)
	copy(s2, s1)
	fmt.Printf("s2: %v\n", s2)
}
*/