/* package main

import (
	"fmt"
)

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() {
	fmt.Println("eat")
}

func (a Animal) sleep() {
	fmt.Println("sleep")
}

type Dog struct {
	Animal
	color string
}

type Cat struct {
	Animal
	ccc string
}

func main() {
	dog := Dog{Animal{"dog", 1}, "red"}
	dog.eat()
	dog.sleep()
	cat := Cat{Animal{"cat", 2}, "ccc"}
	cat.eat()
	cat.sleep()
}
*/