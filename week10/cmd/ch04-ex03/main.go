package main

import (
	"fmt"
	"log"
	"week10/pkg/keyboard"
)

func main() {
	fmt.Print("화씨 온도 입력: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	celsius := (fahrenheit - 32) * 5 / 9

	fmt.Printf("%.2f°F는 %.2f°C입니다.\n", fahrenheit, celsius)
}
