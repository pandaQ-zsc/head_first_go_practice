package main

import (
	"fmt"
)

func main() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}
	fmt.Println(myStruct)
	//当我们调用Printf中的%#v的时候，它将myStruct中的值作为struct 字面量打印
	fmt.Printf("%#v\n", myStruct)
	// 使用电运算符给myStruct赋值
	myStruct.number = 3.14
	fmt.Println(myStruct.number)

	var subscriber struct {
		name   string
		rate   float64
		active bool
	}
	fmt.Println(subscriber)

}
