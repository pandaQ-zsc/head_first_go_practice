package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// GetFloat string -> float64
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	//如果错误不为空， 也就是有错误
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return grade, err
}
func sayHi() {
	fmt.Println("hello")
}

func repeatLine(line string, time int) {
	for i := 0; i < time; i++ {
		fmt.Println(line)
	}
}
func double(number float64) float64 {
	return number * 2
}

func status(grade float64) string {
	if grade < 60.0 {
		return "failing"
	}
	return "passing"
}

func painNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil
}

func main() {
	GetFloat()

	sayHi()
	repeatLine("hello...", 4)
	dozen := double(0.2)
	println(dozen)
	fmt.Println(dozen)
	fmt.Println(status(50))
	//err := errors.New("height can't be negative")
	//fmt.Println(err)
	//log.Fatal(err)
	res, err := painNeeded(-0.233, 0.2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
