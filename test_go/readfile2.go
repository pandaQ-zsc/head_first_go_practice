package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open("data.txt")
	if err != nil {
		return numbers, nil
	}
	// 跟踪我们应该赋值给那个数组索引
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()

	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}

// ...string 是数组
func myFunc(param1 int, param2 ...string) {
	fmt.Println("-----------------")
	fmt.Println(param1)
	fmt.Println(param2)
}

func main() {
	var slice []string
	if len(slice) == 0 {
		slice = append(slice, "first item")
		slice = append(slice, "second item")
	}
	for index, item := range slice {
		fmt.Println(index, "--->", item)
	}
	letter := []string{"a", "b", "c"}
	for i, j := range letter {
		fmt.Println(letter[i])
		fmt.Println(j)

	}
	underlyingString := []string{"a", "b", "ni", "de"}
	slice5 := underlyingString[1:]
	fmt.Println("-----------------")
	for i := range slice5 {
		fmt.Println(slice5[i])
	}
	myFunc(1, "a", "bbbbbbs三十", "c")
}
