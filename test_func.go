/* package main

import "fmt"

func sum(a, b int) (ret int) {
	ret = a + b
	return ret
}

func f1(s []int) {
	s[0] = 100
}

// 可变参数
func f2(args ...int) {
	for _, v := range args {
		fmt.Println("v: ", v)
	}
}

func sayHello(name string) {
	fmt.Printf("Hello,%s\n", name)
}

func f3(name string, f func(string)) {
	f(name)
}

func sub(a, b int) (ret int) {
	ret = a - b
	return ret
}

func cale(op string) func(int, int) int {
	switch op {
	case "+":
		return sum
	case "-":
		return sub
	default:
		return nil
	}
}

// 闭包
func add() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	/* r := sum(1, 2)
	fmt.Printf("r: %v\n", r)
	s := []int{1, 2, 3, 4, 5}
	f1(s)
	fmt.Printf("s: %v\n", s)
	f2(1, 2, 3, 4, 5)
	// 函数类型定义
	type f1 func(int, int) int
	var ff f1
	ff = sum
	r2 := ff(1, 5)
	fmt.Printf("r2: %v\n", r2)
	// 函数作为参数
	f3("zhangsan", sayHello)
	// 函数作为返回值
	f4 := cale("+")
	r3 := f4(1, 3)
	fmt.Printf("r3: %v\n", r3)
	f5 := cale("-")
	r4 := f5(1, 4)
	fmt.Printf("r4: %v\n", r4)
	// 匿名函数
	f6 := func(x, y int) int {
		return x + y
	}
	r4 = f6(1, 7)
	fmt.Printf("r4: %v\n", r4)
	r4 = func(x, y int) int {
		return x + y
	}(1, 8)
	fmt.Printf("r4: %v\n", r4)
	var f7 = add()
	fmt.Printf("f6(1): %v\n", f7(1))
	fmt.Printf("f6(2): %v\n", f7(2))
	f8 := add()
	fmt.Printf("f6(1): %v\n", f8(10))
	fmt.Printf("f6(2): %v\n", f8(20))
}
*/