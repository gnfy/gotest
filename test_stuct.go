/* package main

import "fmt"

// 自定义类型
type Person struct {
	Name string
	Age  int
}

func f1() {
	type Person struct {
		id   int
		name string
	}
	var tom = new(Person)
	tom.id = 1
	tom.name = "tom"
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("tom: %v", *tom)
}

func main() {
	var p1 Person
	fmt.Printf("p1:%v, %T", p1, p1)
	p1.Age = 20
	p1.Name = "Tom"
	fmt.Printf("p1: %v\n", p1)
	var dom struct {
		Name string
		Age  int
	}
	dom.Age = 20
	dom.Name = "dom"
	fmt.Printf("dom: %v\n", dom)

	type p2 struct {
		id   int
		name string
	}
	tom := p2{
		id:   1,
		name: "tom",
	}
	var p3 *p2
	p3 = &tom
	fmt.Printf("p3: %v\n", p3)
	fmt.Printf("p3: %v\n", *p3)
	f1()
}
*/