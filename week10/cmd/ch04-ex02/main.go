package main

import (
	"fmt"
	"log"
	"week10/pkg/keyboard"
)

func main() {
	fmt.Print("점수 입력: ")
	score, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	if score >= 80 {
		fmt.Printf("%.2f점은 합격\n", score)
	} else {
		fmt.Printf("%.2f점은 불합격\n", score)
	}
}
