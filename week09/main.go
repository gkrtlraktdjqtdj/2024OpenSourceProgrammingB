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
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1 // dice 1~6
	fmt.Println(answer)

	fmt.Print("숫자 입력 : ")
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input) // Atoi : sting > integer
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(guess)

	if answer == guess {
		fmt.Println("정답입니다!")
	} else if answer > guess {
		fmt.Println("입력하신 값은 작습니다")
	} else {
		fmt.Println("입력하신 값이 큽니다")
	}
}
