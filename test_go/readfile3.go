package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
		i++
	}
	err = file.Close()
	return numbers, err

}
func main() {
	floatArr, err := getFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, num := range floatArr {
		fmt.Println(num)
	}
}
