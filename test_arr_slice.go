package main

import "fmt"

func main() {
	// 无限数组与切片的区别

	var arr1 = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("arr1: %v\n", arr1)
	// 不能超出数组下标
	// arr1[5] = 10
	arr1[0] = 10
	fmt.Printf("arr1：%v\n", arr1)

	var arr2 = []int{1, 2, 3, 4, 5}
	fmt.Printf("arr2: %v\n", arr2)
	arr2 = append(arr2, 6)
	fmt.Printf("arr2: %v\n", arr2)
}
