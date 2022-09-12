/* package main

import "fmt"

type Person struct {
	name string
}

func (p Person) test() {
	fmt.Printf("test: %v\n", p.name)
}

func (p Person) test2() {
	fmt.Printf("test2: %v\n", p.name)
}

func test3(per Person) {
	per.name = "jack"
}

func test4(per *Person) {
	per.name = "jack2"
}

func (per Person) test5() {
	per.name = "jack3"
}

func (per *Person) test6() {
	per.name = "jack4"
} */

func main() {
	p := Person{
		name: "tom",
	}
	// p.test()
	// p.test2()
	// test3(p)
	// fmt.Printf("p: %v\n", p)
	// p2 := &p
	// test4(p2)
	// fmt.Printf("p2: %v\n", *p2)
	p.test5()
	fmt.Printf("p: %v\n", p)
	p2 := &p
	p2.test6()
	fmt.Printf("p2: %v\n", *p2)
}
