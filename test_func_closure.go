package main

func cale(a int) (func(int) int, func(int) int) {
	add := func(b int) int {
		a += b
		return a
	}
	sub := func(b int) int {
		a -= b
		return a
	}
	return add, sub
}

func main() {
	add, sub := cale(10)
	println(add(1), sub(2))
	add1, sub1 := cale(100)
	println(add1(10), sub1(20))
}
