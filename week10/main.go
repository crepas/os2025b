package main

import "greeting" // week10/pkg/greeting 난 프로그램 파일즈가 아니라 D드라이브에 설치해서 오류남

func main() {
	greeting.Hello() // 패키지에 함수명이 소문자로 시작하면 다른 파일에서 참조하지 못한다. (대문자로 시작해야함!)
	greeting.Hi()
}
