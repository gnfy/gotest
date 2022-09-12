/* package main

import "fmt"

type Flyer interface {
	Fly()
}

type Swimmer interface {
	Swim()
}

// 接口组合
type FlyFish interface {
	Flyer
	Swimmer
}

type Fish struct {
}

func (fish Fish) Fly() {
	fmt.Println("fish fly")
}

func (fish Fish) Swim() {
	fmt.Println("fish swim")
}

func main() {
	var fish FlyFish
	fish = Fish{}
	fish.Fly()
	fish.Swim()
}
*/