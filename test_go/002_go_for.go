package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

func main() {
	fmt.Println(reflect.TypeOf(1))
	//时间
	var now2 = time.Now()
	fmt.Println(now2)
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Date())
	fmt.Println(time.Now().Date())
	//字符串
	//broken := "G# r#cks!"
	//replacer := strings.NewReplacer("#", "o")
	//fmt.Println(replacer.Replace(broken))
	broken := "G# r#cks！"
	replace := strings.NewReplacer("#", "o")
	fixed := replace.Replace(broken)
	fmt.Println(fixed)

	fmt.Println(" Please enter your grade: ")
	read := bufio.NewReader(os.Stdin)
	input, _ := read.ReadString('\n')
	fmt.Println(input)

	s := "\t formerly surrounded by space \n"
	fmt.Println(strings.TrimSpace(s))
	reader1 := bufio.NewReader(os.Stdin)
	input2, err := reader1.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input2)

	//reader:= bufio.NewReader(os.Stdin)
	//input,err:= reader.ReadString('\n')
	// 如果错误不为空， 也就是有错误
	//if err!= nil {
	//	log.Fatal(err)
	//}else {
	//fmt.Println(input)
	//}

	//fmt.Println(reflect.TypeOf(input))
	//grade,err := strconv.ParseFloat(input,64)
	//fmt.Println(reflect.TypeOf(grade))
	//fmt.Println(grade)
	//fmt.Println("ssss")

}
