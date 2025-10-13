package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"strings"
	//"reflect"
	//"time"
)

func main() {
	// var length float64 = 3.2
	// var width int = 2
	// fmt.Println("면적은:", int(length)*width)
	// fmt.Println("length > width?", length > float64(width))
	// fmt.Println("형변환", reflect.TypeOf(int(length)))
	// fmt.Println("원본", reflect.TypeOf(length))

	// var now time.Time = time.Now()
	// var month time.Month = now.Month() // month:=now.Month()
	// var day int = now.Day()
	// fmt.Println(month, day)

	// univ := "G# r#cks!"
	// replacer := strings.NewReplacer("#", "o")
	// fixed := replacer.Replace(univ)
	// fmt.Println(fixed)

	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n') // 리턴을 문자열과 에러로 받음
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i)
}
