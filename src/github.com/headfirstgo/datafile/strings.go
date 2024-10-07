package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GetStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil

	return lines, nil
}
func myMap() {
	lines, err := GetStrings("dataString.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
	var names []string
	var counts []int
	for _, line := range lines {
		flag := false
		for i, name := range names {
			if name == line {
				flag = true
				counts[i] = counts[i] + 1
			}
		}
		if !flag {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}

	for i, name := range names {
		fmt.Printf("%s: %d \n", name, counts[i])
	}
}

func main() {
	lines, err := GetStrings("dataString.txt")
	if err != nil {
		log.Fatal(err)
	}
	var names []string
	var counts []int
	for _, line := range lines {
		flag := false
		for i, name := range names {
			if name == line {
				counts[i] = counts[i] + 1
				flag = true
			}
		}
		if !flag {
			names = append(names, line)
			counts = append(counts, 1)
		}
		fmt.Printf("%s: %d \n", line, counts)
	}

	for i, name := range names {
		fmt.Printf("%s: %d \n", name, counts[i])
	}
}
