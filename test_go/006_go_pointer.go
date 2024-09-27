package main

import (
	"fmt"
)

// 需要注意的是:函数大写是公有的
// 需要注意的是:函数小写是私有的
func double1(amount int) {
	amount *= 2
	fmt.Println(amount)
}
func double2(amount *int) {
	*amount *= 2
	fmt.Println(amount)
}

func negate(myBoolean *bool) {
	//*bool = !*bool
	*myBoolean = !*myBoolean
	//fmt.Println(myBoolean)
}
func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)

	amount := 6
	double1(amount)
	fmt.Println(amount)
	fmt.Println("---------------")
	amount2 := 6
	double2(&amount2)
	fmt.Println(amount2)
	//fmt.Println(amount)
	//var myInt int;
	//var myFloat float64
	//var myBool bool
	//var myString string
	//fmt.Println(reflect.TypeOf(&myInt),reflect.TypeOf(&myFloat),reflect.TypeOf(&myBool),reflect.TypeOf(&myString))
	fmt.Println("=========================")
	//var myFloat float64;
	//var myFloatPointer *float64
	//var myInt int;
	//var myIntPointer *int;
	//myFloatPointer = &myFloat
	//myIntPointer = &myInt
	//fmt.Println(myIntPointer,myFloatPointer)
	fmt.Println("========================")
	//var myBool bool
	//myboolPointer := &myBool
	//fmt.Println(myboolPointer)
	//myInt :=4
	//myPointer := &myInt
	//fmt.Println(myPointer)
	//fmt.Println(*myPointer)
	//
	//myFloat := 3.2
	//myFloatPointer := &myFloat
	//fmt.Println(myFloatPointer)
	//fmt.Println(*myFloatPointer)
	fmt.Println("========================")
	myInt := 4
	myIntPointer := &myInt
	fmt.Println(myInt)
	*myIntPointer = 8
	fmt.Println(myInt)
	fmt.Println(*myIntPointer)

}
