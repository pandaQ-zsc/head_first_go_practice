package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// go install ./src/github.com/headfirstgo/average2
// 执行命令后会安装到
// /Users/xxx/go/bin
// ----------------------
func main() {
	arguments := os.Args[1:]
	var sum float64 = 0
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	log.Printf("Average: %0.2f", sum/float64(len(arguments)))
	sampleCount := float64(len(arguments))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
