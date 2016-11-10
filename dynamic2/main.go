package main

import "fmt"

type emptyInterface interface{}

func foo(intf emptyInterface) {
	// Here, I check what is the underlying type
	switch intf.(type) {
	case int:
		val, _ := intf.(int) // No assertion needed, only convert
		fmt.Println("This is a int type with value ", val)
	case string:
		val, _ := intf.(string)
		fmt.Println("This is a string type with value ", val)
	default:
		fmt.Println("This is an unknown type")
	}
}

func main() {
	ifs := []emptyInterface{int(3), string("abc")}
	for _, v := range ifs {
		foo(v)
	}
}
