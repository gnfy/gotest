/* package main

import "fmt"

type Pet interface {
	eat()
	sleep()
}

type Dog struct {
}
type Cat struct {
}

func (dog Dog) eat() {
	fmt.Println("dog eat")
}

func (dog Dog) sleep() {
	fmt.Println("dog sleep")
}

func (cat Cat) eat() {
	fmt.Println("cat eat")
}
func (cat Cat) sleep() {
	fmt.Println("cat sleep")
}

type Person struct {
}

func (p Person) feed(pet Pet) {
	pet.eat()
	pet.sleep()
}

func main() {
	p := Person{}
	p.feed(Dog{})
	p.feed(Cat{})
}
*/