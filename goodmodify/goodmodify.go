package main

import "fmt"

type Animal interface {
	Speak() string
}

type Cat struct{}

func (c Cat) Speak() string { return "meow" }

type Parrot struct {
	isSmart bool
}

func (p *Parrot) Speak() string { p.isSmart = true; return "Am smart" }

func main() {
	var a Animal
	a = new(Parrot)
	fmt.Println(a.Speak())
}
