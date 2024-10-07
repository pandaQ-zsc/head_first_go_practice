package main

import (
	"fmt"
)

func status2(name string) {
	//字面量创建映射
	grades := map[string]float64{"Al": 22, "Ro": 86.4}
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("%s not found grade recorded \n", name)
	} else if grade < 60 {
		fmt.Printf("%s failed with %.2f\n", name, grade)
	}

}

func main() {
	status2("Al")
	status2("Ro")
	status2("XQ")
	fmt.Println("--------------------")
	//var myMap map[string]int;
	//myMap = make(map[string]int)
	myMap := make(map[string]int)
	ranks := make(map[string]int)
	myMap["one"] = 1
	myMap["gold"] = 2
	myMap["bronze"] = 3
	ranks["gold"] = 1
	ranks["bronze"] = 1
	ranks["silver"] = 1
	for k, v := range myMap {
		fmt.Printf("%s: %d\n", k, v)
	}
	delete(myMap, "gold")
	for k, v := range myMap {
		fmt.Printf("%s: %d\n", k, v)
	}

	fmt.Println("----------------")

	isPrime := make(map[int]bool)
	isPrime[0] = false
	fmt.Println(isPrime[0])

	// 字面量创建映射
	//myMap2 := map[string]float64{"1":2.3,"2":3.4}
	myMap3 := map[string]float64{"1": 2.3, "3": 5.9}
	fmt.Println(myMap3)

	myMap4 := map[string]int{"a": 3, "b": 5}
	var ok bool
	var value int
	value, ok = myMap4["a"]
	fmt.Println(value, ok)

	value, ok = myMap4["c"]
	fmt.Println(value, ok)

	fmt.Println("--------------")
	for k, v := range myMap4 {
		fmt.Printf("%s: %d \n", k, v)
	}
	fmt.Println("--------------")
	emptyMap := map[string]int{}
	fmt.Println(emptyMap)
}
