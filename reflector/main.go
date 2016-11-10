package main

import (
	"fmt"
	"reflect"
)

type emptyInterface interface {
}

// func TypeOf(i interface{}) Type
// func ValueOf(i interface{}) Value

func printDetails(i emptyInterface) {
	fmt.Println("Object has TypeOf ", reflect.TypeOf(i), " and value of ", reflect.ValueOf(i))
}

func main() {
	emptyInterface := []emptyInterface{int(5), float64(5.6), string("Hello World")}
	for _, e := range emptyInterface {
		printDetails(e)
	}
}
