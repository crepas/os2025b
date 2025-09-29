package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var name string = "Kim Inha"
	// var id int = 1000
	// var id int
	// var gpa float32 = 3.99
	// gpa := 3.99 float64(double)

	// name := "kim inha"
	// id := 1000
	// id = 1000

	//zero value
	var f64 float64
	var i16 int16
	var t bool
	var s string
	var i int

	fmt.Println(f64, reflect.TypeOf(f64))
	fmt.Println(i16, reflect.TypeOf(i16))
	fmt.Println(t, reflect.TypeOf(t))
	fmt.Println(s, reflect.TypeOf(s))
	fmt.Println(i, reflect.TypeOf(i))
}
