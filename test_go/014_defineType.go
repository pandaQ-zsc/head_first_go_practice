package main

import "fmt"

// Gallons 和Liters  加仑 和 升
type Gallons float64
type Liters float64

func ToGallon(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

type MyType string

func (m MyType) sayHi() {
	fmt.Println(" hi ", m)
}

type Number int

func (n *Number) Double() {
	*n *= 2
}

// m:接收器（参数） ， MyType:接收器类型
func (m MyType) WithReturn() int {
	return len(m)
}
func main() {
	number := Number(29)
	number.Double()
	fmt.Println(number)

	//方法定义 和 方法调用
	// 先声明一个 value : 接收器
	value := MyType("hello world")
	// 接收器调用方法
	fmt.Println(value.WithReturn())
	value.sayHi()
	fmt.Println("--------------------")
	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)
	fmt.Println(carFuel, busFuel)

	carFuel += ToGallon(Liters(40))
	busFuel += ToLiters(Gallons(10))
	fmt.Println(carFuel, busFuel)

}
