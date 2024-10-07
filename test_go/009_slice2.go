package main

import (
	"fmt"
	"math"
)

func twoInts(first int, second int) {
	fmt.Println(first, second)
}

func myFunc2(param1 int, param2 ...string) {
	fmt.Println("-----------------", param1, param2)
}

func severalString(strings ...string) {
	fmt.Println(strings)
}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}
func maximum(numbers ...float64) float64 {
	maxNum := math.Inf(-1)
	for _, number := range numbers {
		if number > maxNum {
			maxNum = number
		}
	}
	return maxNum
}
func average3(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))

}
func main() {
	myFunc2(1, "hello", "world")
	severalString("aa", "bb")
	severalString("aa")
	severalString()
	fmt.Println(maximum(8, 12, 1, 2, 3, 4, 5))
	fmt.Println(average3(12, 32, 12, 33, 100))
}
