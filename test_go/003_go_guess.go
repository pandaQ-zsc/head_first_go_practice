package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {
	/*
		rand.Intn(n) 返回的随机数范围是[0, n)
	*/

	fmt.Printf("About one-third: %0.2f\n", 1.0/13)

	timeNow := time.Now()
	fmt.Println(timeNow)
	seconds := time.Now().Unix()
	fmt.Println(seconds)
	rand.Seed(seconds)
	//target := rand.Intn(100) + 1
	//fmt.Println(target)

	//  将字符串转化为整型
	parseInt, err := strconv.ParseInt("123", 10, 64)
	//  将字符串转化为浮点型
	//parseFloat, err := strconv.ParseFloat("123", 64)
	//if err != nil {
	//	return
	//}
	fmt.Println(reflect.TypeOf(parseInt))
	fmt.Println(parseInt)

	success := false
	target := rand.Intn(100 + 1)
	fmt.Println("target : ", target)

	if err != nil {
		log.Fatal(err)
	}
	for guessess := 0; guessess < 365; guessess++ {
		fmt.Println("Guessess:", guessess)

		fmt.Println("Guess Game Will Start........")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("你猜测小了")
		} else if guess > target {
			fmt.Println("你猜测大了")
		} else {
			success = true
			fmt.Println("你猜测对了")
			break
		}
		//var x int
		//for x = 1; x < 3; x++ {
		//	fmt.Println(x)
		//}
		//fmt.Println(x)
		//fmt.Println(x)

		//for i := 0; i < 10; i++ {
		//target := rand.Intn(100) + 1
		//fmt.Println(target)
		//}100

	}

	if !success {
		fmt.Println("Game Over ,  It was:", target)
	}
}
