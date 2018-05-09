package main

import (
	"fmt"
)

type Swimmer interface {
	Swim()
}

type Athelete struct{}
type Animal struct{}
type SwimmerImpl struct{}

func (s *Swimmer) Swim() {
	fmt.Println("Swimming")
}

func (a *Athelete) Train() {
	fmt.Println("Training")
}

func (a *Animal) Eat() {
	fmt.Println("Eating")
}

type CompositeSwimmerA struct {
	MyAthelete Athelete
	MySwim     func()
}

type Shark struct {
	Animal
	Swim func()
}

func Swim() {
	fmt.Println("Swimming!")
}

func main() {
	localSwim := Swim
	swimmer := CompositeSwimmerA{
		MySwim: localSwim,
	}
	swimmer.MyAthelete.Train()
	swimmer.MySwim()

	fish := Shark{
		Swim: Swim,
	}

	fish.Eat()
	fish.Swim()
}
