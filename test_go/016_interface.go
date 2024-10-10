package main

import (
	"fmt"
	"head_first_go_practice/src/mypkg"
)

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type Robot string

func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

func (r Robot) MakeSound() {
	fmt.Println("BEEP BOOP")
}

type NoiseMaker interface {
	MakeSound()
	Walk()
}

func play(n NoiseMaker) {
	n.MakeSound()
	n.Walk()
}

type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
}

type Toggleable interface {
	toggle()
}

func main() {
	s := Switch("off")
	//var t Toggleable =&s
	t := &s
	fmt.Println(*t)
	t.toggle()
	fmt.Println(*t)
	t.toggle()
	fmt.Println(*t)

	play(Robot("hello"))

	play(Robot("Toyco"))

	noiseMaker := Robot("1111")
	fmt.Println("==============")
	noiseMaker.Walk()
	//....................

	value := mypkg.MyType(5)
	value.MethodWithParameter(127.3)
	value.MethodWithoutParameters()
	fmt.Println(value.MethodWithReturnValue())
}
