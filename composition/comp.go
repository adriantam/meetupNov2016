package main

type Walker interface {
	Walk() error
}

type Animal interface {
	Speak() error
}

type Human interface {
	Walker
	Animal
}

type Programmer struct {
}

func (p Programmer) Walk() error {
	fmt.Println("Programmer walks slowly")
}

func (p Programmer) Speak() error {
	fmt.Println("Programmer speaks GoLang")
}
