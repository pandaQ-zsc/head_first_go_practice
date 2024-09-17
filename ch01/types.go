package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello, Go!"))
	fmt.Println(reflect.TypeOf(42.7))
	fmt.Println(reflect.TypeOf(42.0))
	fmt.Println(reflect.TypeOf(42.7) == reflect.TypeOf(42.0))
	fmt.Println(reflect.TypeOf(42) == reflect.TypeOf(42.0))
	fmt.Println(reflect.TypeOf(42) == reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(42) == reflect.TypeOf(42.7))
	fmt.Println()

}
