package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 34}
	fmt.Println(mySlice)

	letter := []string{"a", "b", "c"}
	for i := 0; i < len(letter); i++ {
		fmt.Println(letter[i])
	}
	fmt.Println(letter)
	aa := letter[2:]
	fmt.Println(aa)

}
