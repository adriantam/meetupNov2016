package main

import "sort"

type mystruct struct {
	name string
}

// Notice it does not implement method
// Len(), Swap() as well as Less()

func main() {
	mystructs := []mystruct{
		mystruct{name: "abc"},
		mystruct{name: "efg"},
	}
	sort.Sort(mystructs)
}
