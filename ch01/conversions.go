package main

import "fmt"

func main() {
	var length float64 = 1.2
	var width int = 1
	fmt.Println("length is", length)
	fmt.Println("Area is", length*float64(width))
	fmt.Println("length > width?", length > float64(width))
	fmt.Println("width is", width)
	fmt.Println("value of length * width is", length*float64(width))
	fmt.Println("value of width * length is", float64(width)*length)

}
