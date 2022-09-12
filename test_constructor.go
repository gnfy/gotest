package main

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("name is empty")
	}
	if age < 0 {
		return nil, fmt.Errorf("age is negative")
	}
	return &Person{name: name, age: age}, nil
}

func main() {
	per, err := NewPerson("tom", -18)
	if err == nil {
		fmt.Printf("per: %v", *per)
	}
	if err != nil {
		fmt.Printf("err: %v", err)
	}
}
