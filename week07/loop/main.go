package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

	// r := bufio.NewReader(os.Stdin)
	// i, err := r.ReadString('\n') // 리턴을 문자열과 에러로 받음
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(i)

	// var int int = 99 // 예약어 변수명으로 사용하면 안된다
	// fmt.Println(int)

	dice := rand.Intn(6) + 1
	fmt.Println(dice)

	fmt.Print("Enter a score: ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)                // 문자열 주위에 붙은 공란 및 탭 키 등 제거
	score, err := strconv.ParseFloat(i, 64) // 정리된 문자열을 실수타입으로 변환
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if score >= 60 {
		status = "Pass"
	} else {
		status = "Fail"
	}
	fmt.Println(score, status)
}
