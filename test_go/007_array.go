package main

var counter [3]int

func average(arr []float64) float64 {
	var sum float64 = 0
	for i := range arr {
		sum += arr[i]
	}
	return sum / float64(len(arr))
}

func main() {

	//var arr1 = []float64{1, 2, 3, 4}
	//fmt.Println(average(arr1))

	//	counter[0] = 1
	//	counter[1] = 2
	//	counter[2] = 3
	//	println(counter[0])
	//	println(counter[1])
	//	println(counter[2])
	//fmt.Println("====================")
	//	var arr3 = []int{3,1,1,1,1}
	//	for i := range arr3 {
	//		fmt.Println(arr3[i])
	//	}
	//	text := []string{
	//		"ssdsdsd",
	//		"hahahah",
	//		"xq",
	//	}
	//	for i := range text {
	//		fmt.Println(i)
	//	}
	//	for i, e := range text {
	//		fmt.Println(i,e)
	//	}

	//notes  := [3]string {"do","re","mi"}
	//var primes   = [5]int{2,3,5,7}
	//fmt.Println(notes)
	//fmt.Println(primes)
	//fmt.Printf("%#v\n",notes)
	//fmt.Printf("%#v\n",primes)
}
