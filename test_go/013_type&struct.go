package main

import (
	"fmt"
	"head_first_go_practice/src/github.com/headfirstgo/magzine"
)

// 自定义类型
type car struct {
	name     string
	topSpeed int
}
type part struct {
	description string
	count       int
}
type subscriber struct {
	name   string
	rate   float64
	active bool
}

func showInfo(p *part) {
	fmt.Println("Description: ", p.description)
	fmt.Println("Count: ", p.count)
}
func applyDiscount(s *subscriber) {
	//*s.rate = 3.99
	s.rate = 3.99
}
func main() {
	//方式1
	//var sub1 magzine.Subscriber
	//sub1.Rate = 4.99
	//方式2
	sub1 := magzine.Subscriber{Name: "Aman", Rate: 4.99, Active: true}
	fmt.Println(sub1)
	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("Name: ", porsche.name)
	fmt.Println("Top speed:", porsche.topSpeed)
	var bolts part

	// 六角螺栓
	bolts.description = " Hex bolts"
	bolts.count = 24
	fmt.Println("Description: ", bolts.description)
	fmt.Println("Count: ", bolts.count)
	showInfo(&bolts)
	var s subscriber
	applyDiscount(&s)
	fmt.Println(s)

	fmt.Println("--------------")
	// employee 变量
	var employee magzine.Employee
	employee.Name = "xq"
	employee.Salary = 100000
	//employee := magzine.Employee{
	//	Name:   "John",
	//	Salary: 50000,
	//}
	fmt.Println(employee)
	// Subscriber 嵌入struct Street
	subscriber := magzine.Subscriber{
		Name: "John",
		Rate: 4.99,
		//Street:"123Main"
	}
	subscriber.Street = "123 Main"
	fmt.Println(subscriber)

}
