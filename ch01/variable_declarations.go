package main

import (
	"fmt"
)

func main() {
	var (
		quantity      int
		length, width float64
		customerName  string
	)

	quantity = 4
	length, width = 1.2, 2.4
	customerName = "Damon Cole"
	customerName = "Demo Cole"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")
}
