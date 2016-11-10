package main

import "fmt"

func main() {
	var b interface{}
	b = int(3)
	b = float64(9.99)
	b = []string{"Hello", "World"}
	b = new(bool)
	fmt.Println(b)
}
