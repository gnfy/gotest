/* package main

import "fmt"

type Dog struct {
	name  string
	age   int
	color string
}

type Person struct {
	dog  Dog
	id   int
	name string
}

func showPerson(p Person) {
	p.id = 10
	p.name = "tom"
	fmt.Printf("p: %v", p)
}

func showPerson2(p *Person) {
	p.id = 10
	p.name = "tom"
	fmt.Printf("p: %v", p)
}

func main() {
	dog := Dog{
		name:  "dog1",
		age:   10,
		color: "red",
	}
	tom := Person{
		dog:  dog,
		id:   1,
		name: "tom2",
	}
	fmt.Printf("tom: %v\n", tom)
	showPerson(tom)
	p := &tom
	showPerson2(p)
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("p: %v\n", *p)
}
*/