package mypkg

import "fmt"

type myInterface interface {
	methodWithoutParameters()
	methodWithParameter(float64)
	methodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithReturnValue() string {
	return "methodWithReturnValue"
}
func (m MyType) MethodWithoutParameters() {
	fmt.Println("methodWithoutParameters")
}
func (m MyType) MethodWithParameter(f float64) {
	fmt.Println("methodWithParameter")
}

func (m MyType) MethodNoInInterface() {
	fmt.Println("MethodNoInInterface called")
}
