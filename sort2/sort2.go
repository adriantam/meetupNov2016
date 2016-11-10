package main

import "sort"
import "fmt"

type moo struct {
	key   string
	value int
}

type foo []moo

func (f foo) Len() int {
	return len(f)
}
func (f foo) Less(i, j int) bool {
	return f[i].value < f[j].value
}
func (f foo) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func main() {
	f := foo{moo{key: "Y", value: 5}, moo{key: "X", value: 4}, moo{key: "Z", value: 3}}
	sort.Sort(f)
	fmt.Println(f)
}
