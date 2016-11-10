package main

import "sort"
import "fmt"

type foo []string

func (f foo) Len() int {
	return len(f)
}
func (f foo) Less(i, j int) bool {
	return f[i] < f[j]
}
func (f foo) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func main() {
	f := foo{"Y", "X", "Z"}
	sort.Sort(f)
	fmt.Println(f)
}
