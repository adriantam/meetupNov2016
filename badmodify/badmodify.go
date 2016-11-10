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

// Notice parrot will change member variable isSmart
func (p *Parrot) Speak() string { p.isSmart = true; return "Am smart" }

func main() {
	var a Animal
	a = Parrot{} // Will this work?
	fmt.Println(a.Speak())
}
