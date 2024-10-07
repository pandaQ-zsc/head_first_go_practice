package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	var sum float64 = 0
	for _, arguments := range arguments {
		number, err := strconv.ParseFloat(arguments, 64)
		if err != nil {
			log.Fatal(err)
		}

	}
}
