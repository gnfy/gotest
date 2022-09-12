/* package main

import "fmt"

type USBer interface {
	read()
	write()
}

type Computer struct {
	name string
}

func (c Computer) read() {
	fmt.Printf("c.name:%v\n", c.name)
	fmt.Println("read")
}

func (c Computer) write() {
	fmt.Printf("c.name:%v\n", c.name)
	fmt.Println("write")
}

type Pet interface {
	eat()
}
type Dog struct {
}
type Cat struct {
}

func (dog Dog) eat() {
	fmt.Println("dog eat")
}
func (cat Cat) eat() {
	fmt.Println("cat eat")
}
func main() {
	c := Computer{
		name: "lenovo",
	}
	c.read()
	c.write()
	dog := Dog{}
	cat := Cat{}
	dog.eat()
	cat.eat()
	var pet Pet
	pet = Dog{}
	pet.eat()
	pet = Cat{}
	pet.eat()
}
*/