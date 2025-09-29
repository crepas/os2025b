package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var name string = "Kim Inha"
	// var id int = 1000

	name := "kim inha"
	id := 1000

	fmt.Println(name, reflect.TypeOf(name))
	fmt.Println(id, reflect.TypeOf(id))
}
