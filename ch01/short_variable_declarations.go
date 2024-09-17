package main

import "fmt"

func main() {

	price := 100
	taxRate := 0.08
	tax := float64(price) * taxRate
	fmt.Println(tax)
	//fmt.Println(tax)
	availableFunds := 120
	total := float64(price) + tax
	fmt.Println(availableFunds, "dollars available")
	fmt.Println("within budget?", total <= float64(availableFunds))
	fmt.Println("")
	fmt.Println()
	//  源代码文件 进行编译， 将其转化为cpu可以执行的二进制文件
}
