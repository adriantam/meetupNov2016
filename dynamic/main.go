package main

import "fmt"

type emptyInterface interface{}

func foo(intf emptyInterface) {
	var i int
	// Here, I assert that intf is of type int
	// and convert the value into variable i
	i, ok := intf.(int)
	if ok {
		fmt.Println("This is a int type with value ", i)
	} else {
		fmt.Println("This is not an int type")
	}
}

func main() {
	ifs := []emptyInterface{int(3), string("abc")}
	for _, v := range ifs {
		foo(v)
	}
}
