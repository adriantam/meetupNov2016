package main

import "fmt"

type Equaler interface {
	Equal(Equaler) bool
}

type T int

func (t T) Equal(u Equaler) bool { return t == u.(T) } // now satisfies Equaler

func main() {
	var b Equaler
	var c Equaler
	b = T(3)
	c = T(4)
	fmt.Println(b.Equal(c))
}
