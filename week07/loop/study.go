package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	s := time.Now().Unix()
	rand.Seed(s)
	num := rand.Intn(100) + 1
	fmt.Println("하이 앤 로우 게임 시작 1~100 사이의 숫자를 맞춰보세요.")

	success := false

	for i := 1; i <= 5; i++ {
		fmt.Printf("%d번째 시도 : ", i)
		r := bufio.NewReader(os.Stdin)
		input, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		g, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if g < num {
			fmt.Println("하이")
		} else if g > num {
			fmt.Println("로우")
		} else {
			fmt.Println("정답입니다!")
			success = true
			break
		}
	}

	if !success {
		fmt.Printf("정답은 %d였습니다.\n", num)
	} else {
		fmt.Println("게임에 참여해 주셔서 감사합니다!")
	}
}
