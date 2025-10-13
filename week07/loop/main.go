package main

import (
	"fmt"
	"strings"

	//"reflect"
	"time"
)

func main() {
	// var length float64 = 3.2
	// var width int = 2
	// fmt.Println("면적은:", int(length)*width)
	// fmt.Println("length > width?", length > float64(width))
	// fmt.Println("형변환", reflect.TypeOf(int(length)))
	// fmt.Println("원본", reflect.TypeOf(length))
	var now time.Time = time.Now()
	var month time.Month = now.Month() // month:=now.Month()
	var day int = now.Day()
	fmt.Println(month, day)

	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
