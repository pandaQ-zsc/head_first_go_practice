package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getGloat(fileName string) ([]float64, error) {
	//file,err:=os.Open("data.txt")
	file, err := os.Open(fileName)
	var numbers []float64
	if err != nil {
		log.Fatal()
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, number)
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

func getFloat2(filename string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		//err: 不用“:=”重新声明 numbers := append(numbers,number)
		numbers = append(numbers, number)

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
func main() {
	fmt.Println("------------")
	fmt.Println(os.Args)
	fmt.Println(getGloat("data.txt"))
	notes := make([]string, 8)
	notes[0] = "a"
	notes[1] = "a"
	notes[2] = "a"
	fmt.Println(notes)
	fmt.Println("=============================")
	mySlice := []int{1, 2, 34}
	fmt.Println(mySlice)
	letter := []string{"a", "b", "c"}
	for i := 0; i < len(letter); i++ {
		fmt.Println(letter[i])
	}
	fmt.Println(letter)
	aa := letter[2:]
	fmt.Println(aa)
	underlyArray := [5]string{"aaaa", "p", "t", "d", "bb"}
	fmt.Println(underlyArray)
	arrSlice := underlyArray[3:]
	fmt.Println(arrSlice)
	underlyArray[3] = "dd"
	fmt.Println(arrSlice)
	fmt.Println("============")
	fmt.Println(underlyArray)
	arrSlice[1] = "cc"
	fmt.Println(arrSlice)
	fmt.Println(underlyArray)
	array1 := [5]int{1, 3, 4, 5, 3}
	slice1 := array1[1:3]
	fmt.Println(slice1)

	array2 := []string{"a", "c", "d", "e", "f"}
	fmt.Println(array2[2:])
	array2 = append(array2, "c")
	fmt.Println(array2[2:])
	fmt.Println(len(slice1))
	var slice []string
	if len(slice) == 0 {
		slice = append(slice, "first item")
	}
	fmt.Println("嘿嘿嘿")

}
