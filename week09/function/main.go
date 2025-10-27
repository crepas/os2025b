package main

import (
	"fmt"
)

func swap(first int, second int) {
	temp := first
	first = second
	second = temp
	fmt.Println(first, second)
}

func main() {
	//fmt.Println(math.Sqrt(16))
	a, b := 10, 25
	fmt.Println(a, b)
	swap(a, b)
	fmt.Println(a, b)
}
