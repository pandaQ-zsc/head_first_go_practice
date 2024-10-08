package magzine

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	// 嵌入
	Address
}
type Employee struct {
	Name   string
	Salary float64
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode int
}
