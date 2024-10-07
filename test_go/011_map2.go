package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func getStrings(fileName string) ([]string, error) {
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

func main() {
	counts := make(map[string]int)
	counts["aa"]++
	counts["aa"]++
	counts["cc"]++
	counts["dd"]++
	fmt.Println(counts)
	lines, err := getStrings("dataString.txt")
	if err != nil {
		log.Fatal(err)
	}
	lineMap := make(map[string]int)
	for _, line := range lines {
		lineMap[line]++
	}
	fmt.Println(lineMap)
	grades := map[string]float64{"A": 74.2, "B": 86.5, "C": 59.7, "D": 22.2}
	for k, v := range grades {
		fmt.Printf("%s: %.2f\n", k, v)
		if v < 60 {
			fmt.Println("low")
		} else if v > 60 {
			fmt.Println("high")
		}
	}
	fmt.Println("Class roster:")
	for k, name := range grades {
		fmt.Println(k, name)
	}
	var names []string
	for name, _ := range grades {
		names = append(names, name)
	}
	// 使用sort.Strings() 按照字母表顺序给切片排序
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s : %0.1f%%\n", name, grades[name])
	}
}
