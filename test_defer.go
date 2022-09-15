/* package main

import "fmt"

func f1() int {
	x := 1
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// x作用域是整个函数，defer语句中的x是一个新的变量，与函数中的x不是同一个变量
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	fmt.Println("Hello, playground")
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	a := f1()
	fmt.Printf("a: %v\n", a)
	b := f2()
	fmt.Printf("b: %v\n", b)
	c := f3()
	fmt.Printf("c: %v\n", c)
	d := f4()
	fmt.Printf("d: %v\n", d)
}
*/