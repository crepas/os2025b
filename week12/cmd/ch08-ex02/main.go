package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func main() {
	var s1 magazine.Subscriber
	var e1 magazine.Employee
	s1.Name = "최인화"
	e1.Name = "이인하"
	e1.Salary = 10000000
	e1.Address.City = "인천"
	s1.Address.City = "서울"
	fmt.Println(e1.Name, e1.Salary, e1.Address.City)
	fmt.Println(s1.Name, s1.Address.City)
}
