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
	fmt.Println("하이 앤 로우 게임 시작 1~100 사이의 숫자를 맞춰보세요.")
	s := time.Now().Unix()
	rand.Seed(s)
	num := rand.Intn(100) + 1

	success := false
	for i := 1; i <= 5; i++ {
		fmt.Println(i, "번째 시도")
		input := bufio.NewReader(os.Stdin)
		i, err := input.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		i := strings.TrimSpace(i)
		i, err := strconv.Atoi(i)
	}
}
